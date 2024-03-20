package core

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Kinto addresses mainnet
/*var (
	aaEntryPointEnvAddress = common.HexToAddress("0x2843C269D2a64eCfA63548E8B3Fc0FD23B7F70cb")
	kintoIdEnvAddress      = common.HexToAddress("0xf369f78E3A0492CC4e96a90dae0728A38498e9c7")
	walletFactoryAddress   = common.HexToAddress("0x8a4720488CA32f1223ccFE5A087e250fE3BC5D75")
	paymasterAddress       = common.HexToAddress("0x1842a4EFf3eFd24c50B63c3CF89cECEe245Fc2bd")
	appRegistryAddress     = common.HexToAddress("0x5A2b641b84b0230C8e75F55d5afd27f4Dbd59d5b")
	upgradeExecutor        = common.HexToAddress("0x88e03D41a6EAA9A0B93B0e2d6F1B34619cC4319b")
	l2ERC20GatewayAddress  = common.HexToAddress("0x87799989341A07F495287B1433eea98398FD73aA")
)*/

// Kinto addresses devnet
var (
	aaEntryPointEnvAddress = common.HexToAddress("0x40Ec0101AEA7A1CC550E445e641AB59dec6daff7")
	kintoIdEnvAddress      = common.HexToAddress("0x6d2f5f6f0E633c10b6AC4610f09F0392c65128C2")
	walletFactoryAddress   = common.HexToAddress("0xa3F85Ea46fA7f1008c0061F80c433231f3833700")
	paymasterAddress       = common.HexToAddress("0xe17E0001A8Df51F8778020c021C11dA76b3dAe2D")
	appRegistryAddress     = common.HexToAddress("0xE3BF35068FaA931259E3F200Ce567da5EC8CC18f")
	upgradeExecutor        = common.HexToAddress("0x88e03D41a6EAA9A0B93B0e2d6F1B34619cC4319b")
	l2ERC20GatewayAddress  = common.HexToAddress("0x87799989341A07F495287B1433eea98398FD73aA")
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
	l2ERC20GatewayAddress:  true, // l2ERC20GatewayAddress
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
