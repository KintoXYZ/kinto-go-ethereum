package core

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

var hardfork6KintoAddresses = map[common.Address]bool{
	aaEntryPointEnvAddress:      true,  // aaEntryPointEnvAddress
	kintoIdEnvAddress:           true,  // kintoIdEnvAddress
	walletFactoryAddress:        true,  // walletFactoryAddress
	paymasterAddress:            true,  // paymasterAddress
	appRegistryAddress:          true,  // appRegistryAddress
	upgradeExecutor:             true,  // upgradeExecutor
	customGatewayAddress:        true,  // customGatewayAddress
	gatewayRouterAddress:        true,  // gatewayRouterAddress
	standardGatewayAddress:      true,  // standardGatewayAddress
	wethGateWayAddress:          true,  // wethGateWayAddress
	bundleBulker:                true,  // bundleBulker
	arbRetrayableTx:             true,  // arbRetrayableTx
	socket:                      false, // socket
	socketExecutionManager:      false, // socketExecutionManager
	socketTransmitManager:       false, // socketTransmitManager
	socketFastSwitchboard:       false, // socketFastSwitchboard
	socketOptimisticSwitchboard: false, // socketOptimisticSwitchboard
	socketBatcher:               false, // socketBatcher
	socketSimulator:             false, // socketSimulator
	socketSimulatorUtils:        false, // socketSimulatorUtils
	socketSwitchboardSimulator:  false, // socketSwitchboardSimulator
	socketCapacitorSimulator:    false, // socketCapacitorSimulator
	create2Factory:              false, // create2Factory
	aaEntryPointEnvAddressV7:    true,  // aaEntryPointEnvAddressv7
}

var ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

func enforceHardForkSixRules(st *StateTransition) error {
	msg := st.msg

	if msg.TxRunMode == MessageGasEstimationMode {
		return nil // allow gas estimation
	}

	destination := msg.To
	origin := msg.From

	if destination == nil {
		destination = &ZeroAddress
	}

	allowed, err := isContractCallAllowedFromEOA(st, origin, *destination)

	if allowed && err == nil {
		return nil
	}

	functionSelector := extractFunctionSelector(msg.Data)

	if *destination == ZeroAddress {
		return fmt.Errorf("%w: %v EOAs can't create contracts directly, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if _, ok := hardfork6KintoAddresses[*destination]; !ok {
		return fmt.Errorf("%w: Transaction to address %v is not permitted", ErrKintoNotAllowed, destination.Hex())
	}

	if isEntryPointAddress(*destination) && isEntryPointWithdraw(functionSelector) {
		addressBytes := msg.Data[functionSelectorSize+addressOffset : functionSelectorSize+fullWordSize]
		paramAddress := common.BytesToAddress(addressBytes)

		if msg.From != paramAddress {
			return fmt.Errorf("%w: %v is trying to withdrawTo/withdrawStake from EntryPoint to a param different than the sender, %v", ErrKintoNotAllowed, msg.From.Hex(), paramAddress)
		}
	}

	if isHandleOps(*destination, functionSelector) {
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

	if isEntryPointAddress(*destination) && hardForkSixForbiddenEPFunctions(functionSelector) {
		return fmt.Errorf("%w: %v EntryPoint depositTo, HandleAggregatedOps and fallback functions are not allowed , %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if *destination == paymasterAddress && paymasterFunctionNotAllowed(functionSelector) { //ENTRYPOINT PAYMASTER RULES
		return fmt.Errorf("%w: %v SponsorPaymaster withDrawTo() and deposit() are not allowed , %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	return nil
}

func isContractCallAllowedFromEOA(st *StateTransition, from, to common.Address) (bool, error) {
	fmt.Printf("****Checking if contract call is allowed from EOA\n")
	//log from and to
	fmt.Printf("****From: %v\n", from)
	fmt.Printf("****To: %v\n", to)
	// Define the ABI
	const abiJSON = `[{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"}],"name":"isContractCallAllowedFromEOA","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return false, fmt.Errorf("error parsing ABI: %v", err)
	}

	input, err := parsedABI.Pack("isContractCallAllowedFromEOA", from, to)
	if err != nil {
		return false, fmt.Errorf("error packing function call: %v", err)
	}

	ret, _, err := st.evm.Call(vm.AccountRef(from), appRegistryAddress, input, uint64(100000), uint256.NewInt(0))
	if err != nil {
		return false, fmt.Errorf("error executing contract call: %v", err)
	}

	if len(ret) == 0 {
		return false, fmt.Errorf("empty result from contract call")
	}

	var result bool
	err = parsedABI.UnpackIntoInterface(&result, "isContractCallAllowedFromEOA", ret)
	if err != nil {
		return false, fmt.Errorf("error unpacking result: %v", err)
	}
	//log result
	fmt.Printf("****Result: %v\n", result)

	return result, nil
}

func isEntryPointAddress(address common.Address) bool {
	return address == aaEntryPointEnvAddress || address == aaEntryPointEnvAddressV7
}

func isHandleOps(address common.Address, functionSelector string) bool {
	return isEntryPointAddress(address) && (functionSelector == functionSelectorEPHandleOps || functionSelector == functionSelectorEPHandleOpsV7)
}

func hardForkSixForbiddenEPFunctions(functionSelector string) bool {
	return (functionSelector == functionSelectorEmpty ||
		functionSelector == functionSelectorEPDeposit ||
		functionSelector == functionSelectorEPHandleAggregatedOps ||
		functionSelector == functionSelectorEPHandleOpsV7)
}
