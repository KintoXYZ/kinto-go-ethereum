package core

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Kinto addresses mainnet
/*
var (
	aaEntryPointEnvAddress = common.HexToAddress("0x2843C269D2a64eCfA63548E8B3Fc0FD23B7F70cb")
	kintoIdEnvAddress      = common.HexToAddress("0xf369f78E3A0492CC4e96a90dae0728A38498e9c7")
	walletFactoryAddress   = common.HexToAddress("0x8a4720488CA32f1223ccFE5A087e250fE3BC5D75")
	paymasterAddress       = common.HexToAddress("0x1842a4EFf3eFd24c50B63c3CF89cECEe245Fc2bd")
	appRegistryAddress     = common.HexToAddress("0x5A2b641b84b0230C8e75F55d5afd27f4Dbd59d5b")
	upgradeExecutor        = common.HexToAddress("0x88e03D41a6EAA9A0B93B0e2d6F1B34619cC4319b")
	customGatewayAddress   = common.HexToAddress("0x06FcD8264caF5c28D86eb4630c20004aa1faAaA8")
	gatewayRouterAddress   = common.HexToAddress("0x340487b92808B84c2bd97C87B590EE81267E04a7")
	standardGatewayAddress = common.HexToAddress("0x87799989341A07F495287B1433eea98398FD73aA")
	wethGateWayAddress     = common.HexToAddress("0xd563ECBDF90EBA783d0a218EFf158C1263ad02BE")
)
*/

// Kinto addresses devnet
var (
	aaEntryPointEnvAddress = common.HexToAddress("0xEeb65A06722E6B7141114980Fff7d86CCB14F435")
	kintoIdEnvAddress      = common.HexToAddress("0xd7Fa9143481d9c48DF79Bb042A6A7a51C99112B6")
	walletFactoryAddress   = common.HexToAddress("0xB6816E20AfC8412b7D6eD491F0c41317315c29D3")
	paymasterAddress       = common.HexToAddress("0x29C157fb553D9EAD78e5084F74E02F2ACEbE6770")
	appRegistryAddress     = common.HexToAddress("0xF2c5B9400a562c6429db9f015eD705C1CA8458A9")
	upgradeExecutor        = common.HexToAddress("0x6B0d3F40DeD9720938DB274f752F1e11532c2640")
	customGatewayAddress   = common.HexToAddress("0x094F8C3eA1b5671dd19E15eCD93C80d2A33fCA99")
	gatewayRouterAddress   = common.HexToAddress("0xf3AC740Fcc64eEd76dFaE663807749189A332d54")
	standardGatewayAddress = common.HexToAddress("0x6A8d32c495df943212B7788114e41103047150a5")
	wethGateWayAddress     = common.HexToAddress("0x79B47F0695608aD8dc90E400a3E123b02eB72D24")
)

// Kinto-specific constants for function selectors
const (
	functionSelectorEPWithdrawTo          = "205c2878"
	functionSelectorEPWithdrawStake       = "c23a5cea"
	functionSelectorEPHandleOps           = "1fad948c"
	functionSelectorEPHandleAggregatedOps = "4b1d7cf5"
	functionSelectorSPWithdrawTo          = "205c2878"
	functionSelectorSPDeposit             = "d0e30db0"
	functionSelectorEmpty                 = "" //Hardfork2 start
	functionSelectorEPDeposit             = "b760faf9"
)

const (
	functionSelectorSize = 4  // size of the function selector
	addressOffset        = 12 // offset to skip leading zeros in a 32-byte word to get to the 20-byte address
	fullWordSize         = 32 // size of a full 32-byte word, standard in Ethereum for holding a word
	beneficiaryOffset    = 32 // offset to skip the first 32 bytes of the data (function selector) to get to the beneficiary address
)

// Valid Kinto addresses before the hardfork
var originalKintoAddresses = map[common.Address]bool{
	aaEntryPointEnvAddress: true, // aaEntryPointEnvAddress
	kintoIdEnvAddress:      true, // kintoIdEnvAddress
	walletFactoryAddress:   true, // walletFactoryAddress
	paymasterAddress:       true, // paymasterAddress
}

// Valid Kinto addresses after the hardfork
var hardfork1KintoAddresses = map[common.Address]bool{
	aaEntryPointEnvAddress: true, // aaEntryPointEnvAddress
	kintoIdEnvAddress:      true, // kintoIdEnvAddress
	walletFactoryAddress:   true, // walletFactoryAddress
	paymasterAddress:       true, // paymasterAddress
	appRegistryAddress:     true, // appRegistryAddress
}

// Valid Kinto addresses after the hardfork #2
var hardfork2KintoAddresses = map[common.Address]bool{
	aaEntryPointEnvAddress: true, // aaEntryPointEnvAddress
	kintoIdEnvAddress:      true, // kintoIdEnvAddress
	walletFactoryAddress:   true, // walletFactoryAddress
	paymasterAddress:       true, // paymasterAddress
	appRegistryAddress:     true, // appRegistryAddress
	upgradeExecutor:        true, // upgradeExecutor
}

var hardfork3KintoAddresses = map[common.Address]bool{
	aaEntryPointEnvAddress: true, // aaEntryPointEnvAddress
	kintoIdEnvAddress:      true, // kintoIdEnvAddress
	walletFactoryAddress:   true, // walletFactoryAddress
	paymasterAddress:       true, // paymasterAddress
	appRegistryAddress:     true, // appRegistryAddress
	upgradeExecutor:        true, // upgradeExecutor
	customGatewayAddress:   true, // customGatewayAddress
	gatewayRouterAddress:   true, // gatewayRouterAddress
	standardGatewayAddress: true, // standardGatewayAddress
	wethGateWayAddress:     true, // wethGateWayAddress
}

// enforceKinto decides which set of Kinto rules to apply based on the current block number
func enforceKinto(msg *Message, currentBlockNumber *big.Int) error {
	if msg.TxRunMode == MessageEthcallMode {
		return nil // Allow all calls
	}

	if currentBlockNumber.Cmp(common.KintoRulesBlockStart) > 0 {
		if currentBlockNumber.Cmp(common.KintoHardfork1) <= 0 {
			return enforceOriginalKintoRules(msg) // Original Kinto rules
		} else if currentBlockNumber.Cmp(common.KintoHardfork2) <= 0 {
			return enforceHardForkOneRules(msg) // Rules for the first hard fork
		} else if currentBlockNumber.Cmp(common.KintoHardfork3) <= 0 {
			return enforceHardForkTwoRules(msg) // Rules for the second hard fork
		} else {
			return enforceHardForkThreeRules(msg) //Rules for the third hard fork
		}
	}
	return nil
}

// enforceOriginalKintoRules applies the original Kinto rules
func enforceOriginalKintoRules(msg *Message) error {
	destination := msg.To

	if destination == nil {
		return fmt.Errorf("%w: %v is trying to create a contract directly, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if _, ok := originalKintoAddresses[*destination]; !ok {
		return fmt.Errorf("%w: %v is trying to tx against an invalid address, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	return nil
}

// enforceHardForkOneRules applies the Kinto rules after the first hardfork
func enforceHardForkOneRules(msg *Message) error {
	destination := msg.To
	functionSelector := extractFunctionSelector(msg.Data)

	if destination == nil {
		return fmt.Errorf("%w: %v EOAs can't create contracts directly, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if _, ok := hardfork1KintoAddresses[*destination]; !ok {
		return fmt.Errorf("%w: Transaction to address %v is not permitted", ErrKintoNotAllowed, destination.Hex())
	}

	if *destination == aaEntryPointEnvAddress && isEntryPointWithdraw(functionSelector) {
		addressBytes := msg.Data[functionSelectorSize+addressOffset : functionSelectorSize+fullWordSize]
		paramAddress := common.BytesToAddress(addressBytes)

		if msg.From != paramAddress {
			return fmt.Errorf("%w: %v is trying to withdrawTo/withdrawStake from EntryPoint to a param different than the sender, %v", ErrKintoNotAllowed, msg.From.Hex(), paramAddress)
		}
	}

	if *destination == aaEntryPointEnvAddress && isEntryPointHandleOps(functionSelector) {
		data := msg.Data[functionSelectorSize:]
		if len(data) >= beneficiaryOffset+fullWordSize {
			beneficiaryEncoded := data[beneficiaryOffset : beneficiaryOffset+fullWordSize]
			beneficiaryBytes := beneficiaryEncoded[addressOffset:]
			beneficiaryAddress := common.BytesToAddress(beneficiaryBytes)

			if msg.From != beneficiaryAddress {
				return fmt.Errorf("%w: %v is trying to handleOps/handleAggregatedOps from EntryPoint to a beneficiary different than the sender, %v", ErrKintoNotAllowed, msg.From.Hex(), beneficiaryAddress)
			}
		}
	}

	if *destination == paymasterAddress && paymasterFunctionNotAllowed(functionSelector) { //ENTRYPOINT PAYMASTER RULES
		return fmt.Errorf("%w: %v SponsorPaymaster withDrawTo() and deposit() are not allowed , %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	return nil
}

func enforceHardForkTwoRules(msg *Message) error {
	destination := msg.To
	functionSelector := extractFunctionSelector(msg.Data)

	if destination == nil {
		return fmt.Errorf("%w: %v EOAs can't create contracts directly, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if _, ok := hardfork2KintoAddresses[*destination]; !ok {
		return fmt.Errorf("%w: Transaction to address %v is not permitted", ErrKintoNotAllowed, destination.Hex())
	}

	if *destination == aaEntryPointEnvAddress && isEntryPointWithdraw(functionSelector) {
		addressBytes := msg.Data[functionSelectorSize+addressOffset : functionSelectorSize+fullWordSize]
		paramAddress := common.BytesToAddress(addressBytes)

		if msg.From != paramAddress {
			return fmt.Errorf("%w: %v is trying to withdrawTo/withdrawStake from EntryPoint to a param different than the sender, %v", ErrKintoNotAllowed, msg.From.Hex(), paramAddress)
		}
	}

	if *destination == aaEntryPointEnvAddress && functionSelector == functionSelectorEPHandleOps {
		data := msg.Data[functionSelectorSize:]
		if len(data) >= beneficiaryOffset+fullWordSize {
			beneficiaryEncoded := data[beneficiaryOffset : beneficiaryOffset+fullWordSize]
			beneficiaryBytes := beneficiaryEncoded[addressOffset:]
			beneficiaryAddress := common.BytesToAddress(beneficiaryBytes)

			if msg.From != beneficiaryAddress {
				return fmt.Errorf("%w: %v is trying to handleOps from EntryPoint to a beneficiary different than the sender, %v", ErrKintoNotAllowed, msg.From.Hex(), beneficiaryAddress)
			}
		}
	}

	if *destination == aaEntryPointEnvAddress && hardForkTwoForbiddenEPFunctions(functionSelector) {
		return fmt.Errorf("%w: %v EntryPoint depositTo, HandleAggregatedOps and fallback functions are not allowed , %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if *destination == paymasterAddress && paymasterFunctionNotAllowed(functionSelector) { //ENTRYPOINT PAYMASTER RULES
		return fmt.Errorf("%w: %v SponsorPaymaster withDrawTo() and deposit() are not allowed , %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	return nil
}

func enforceHardForkThreeRules(msg *Message) error {
	destination := msg.To
	functionSelector := extractFunctionSelector(msg.Data)

	if destination == nil {
		return fmt.Errorf("%w: %v EOAs can't create contracts directly, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if _, ok := hardfork3KintoAddresses[*destination]; !ok {
		return fmt.Errorf("%w: Transaction to address %v is not permitted", ErrKintoNotAllowed, destination.Hex())
	}

	if *destination == aaEntryPointEnvAddress && isEntryPointWithdraw(functionSelector) {
		addressBytes := msg.Data[functionSelectorSize+addressOffset : functionSelectorSize+fullWordSize]
		paramAddress := common.BytesToAddress(addressBytes)

		if msg.From != paramAddress {
			return fmt.Errorf("%w: %v is trying to withdrawTo/withdrawStake from EntryPoint to a param different than the sender, %v", ErrKintoNotAllowed, msg.From.Hex(), paramAddress)
		}
	}

	if *destination == aaEntryPointEnvAddress && functionSelector == functionSelectorEPHandleOps {
		data := msg.Data[functionSelectorSize:]
		if len(data) >= beneficiaryOffset+fullWordSize {
			beneficiaryEncoded := data[beneficiaryOffset : beneficiaryOffset+fullWordSize]
			beneficiaryBytes := beneficiaryEncoded[addressOffset:]
			beneficiaryAddress := common.BytesToAddress(beneficiaryBytes)

			if msg.From != beneficiaryAddress {
				return fmt.Errorf("%w: %v is trying to handleOps from EntryPoint to a beneficiary different than the sender, %v", ErrKintoNotAllowed, msg.From.Hex(), beneficiaryAddress)
			}
		}
	}

	if *destination == aaEntryPointEnvAddress && hardForkTwoForbiddenEPFunctions(functionSelector) {
		return fmt.Errorf("%w: %v EntryPoint depositTo, HandleAggregatedOps and fallback functions are not allowed , %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if *destination == paymasterAddress && paymasterFunctionNotAllowed(functionSelector) { //ENTRYPOINT PAYMASTER RULES
		return fmt.Errorf("%w: %v SponsorPaymaster withDrawTo() and deposit() are not allowed , %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	return nil
}

func extractFunctionSelector(data []byte) string {
	if len(data) < functionSelectorSize {
		return ""
	}
	return hex.EncodeToString(data[:functionSelectorSize])
}

func paymasterFunctionNotAllowed(functionSelector string) bool {
	return functionSelector == functionSelectorSPWithdrawTo || functionSelector == functionSelectorSPDeposit
}

func isEntryPointHandleOps(functionSelector string) bool {
	return (functionSelector == functionSelectorEPHandleOps || functionSelector == functionSelectorEPHandleAggregatedOps)
}

func isEntryPointWithdraw(functionSelector string) bool {
	return (functionSelector == functionSelectorEPWithdrawTo || functionSelector == functionSelectorEPWithdrawStake)
}

func hardForkTwoForbiddenEPFunctions(functionSelector string) bool {
	return (functionSelector == functionSelectorEmpty ||
		functionSelector == functionSelectorEPDeposit ||
		functionSelector == functionSelectorEPHandleAggregatedOps)
}
