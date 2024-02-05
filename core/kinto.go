package core

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Kinto addresses
var (
	aaEntryPointEnvAddress = common.HexToAddress("0x351110fC667dA12B5d07AEDaE6e90f17BAF512C0")
	kintoIdEnvAddress      = common.HexToAddress("0xa812c34cB952039934B6e0b86E91F628ce0092aa")
	walletFactoryAddress   = common.HexToAddress("0x2fdECA9826f3dA40E7ebe463Bd0BC8CE5a274752")
	paymasterAddress       = common.HexToAddress("0x6ecDCd6C797Cb1D358eB436935095d0b04949fb9")
	appRegistryAddress     = common.HexToAddress("0x79609fCE4791C3f0067aDEc72DcDB1a89cCbf58F")
)

// Kinto-specific constants for function selectors
const (
	functionSelectorEPWithdrawTo          = "205c2878"
	functionSelectorEPWithdrawStake       = "c23a5cea"
	functionSelectorEPHandleOps           = "1fad948c"
	functionSelectorEPHandleAggregatedOps = "4b1d7cf5"
	functionSelectorSPWithdrawTo          = "205c2878"
	functionSelectorSPDeposit             = "d0e30db0"
)

const (
	functionSelectorSize = 4  // size of the function selector
	addressOffset        = 12 // offset to skip leading zeros in a 32-byte word to get to the 20-byte address
	fullWordSize         = 32 // size of a full 32-byte word, standard in Ethereum for holding a word
	beneficiaryOffset    = 32 // offset to skip the first 32 bytes of the data (function selector) to get to the beneficiary address
)

// Block numbers for Kinto rule changes
var (
	KintoRulesBlockStart = big.NewInt(100)
	KintoHardfork1       = big.NewInt(110)
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

// enforceKinto decides which set of Kinto rules to apply based on the current block number
func enforceKinto(msg *Message, currentBlockNumber *big.Int) error {
	if currentBlockNumber.Cmp(KintoRulesBlockStart) > 0 {
		if currentBlockNumber.Cmp(KintoHardfork1) <= 0 {
			if err := enforceOriginalKintoRules(msg); err != nil {
				return err
			}
		} else {
			if err := enforceHardForkOneRules(msg); err != nil {
				return err
			}
		}
	}
	return nil
}

// enforceOriginalKintoRules applies the original Kinto rules
func enforceOriginalKintoRules(msg *Message) error {
	log.Warn("****** KINTO ORIGINAL RULES ******")
	destination := msg.To

	if destination == nil {
		return fmt.Errorf("%w: EOAs can't create contracts directly", ErrKintoNotAllowed)
	}

	if _, ok := originalKintoAddresses[*destination]; !ok {
		return fmt.Errorf("%w: Transaction to address %v is not permitted", ErrKintoNotAllowed, destination.Hex())
	}

	return nil
}

// enforceHardForkOneRules applies the Kinto rules after the first hardfork
func enforceHardForkOneRules(msg *Message) error {
	log.Warn("****** KINTO HARDFORK #1 RULES ******")
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
