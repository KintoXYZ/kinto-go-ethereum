package core

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

//var replaceHF5Address = common.HexToAddress("0x4e59b44847b379578588920cA78FbF26c0B4956C")
//var replacedHF5Bytecode = common.FromHex("0x6080604052600436106100225760003560e01c806399a6cddd146101225761004d565b3661004d57604051638dc2b20960e01b81523360048201523460248201526044015b60405180910390fd5b6040516313289ea360e31b815233600482015260009036906060906001600160a01b037f000000000000000000000000f369f78e3a0492cc4e96a90dae0728a38498e9c71690639944f51890602401602060405180830381865afa1580156100b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100dd9190610172565b6100fc57604051630ab529c360e21b8152336004820152602401610044565b36601f1901600081602082378035828234f5915081610119578081fd5b8181526014600cf35b34801561012e57600080fd5b506101567f000000000000000000000000f369f78e3a0492cc4e96a90dae0728a38498e9c781565b6040516001600160a01b03909116815260200160405180910390f35b60006020828403121561018457600080fd5b8151801515811461019457600080fd5b939250505056fea2646970667358221220a35916f465fc2e7b19efa5b0d984bc014269d1d8cbad39f684ed4ecc1dd45e5f64736f6c63430008180033")

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

func enforceHardForkSixRules(st *StateTransition) error {
	msg := st.msg

	if msg.TxRunMode == MessageGasEstimationMode {
		return nil // allow gas estimation
	}

	destination := msg.To
	origin := msg.From

	allowed, err := isContractCallAllowedFromEOA(st, origin, *destination)

	if allowed && err == nil {
		return nil
	}

	functionSelector := extractFunctionSelector(msg.Data)

	if destination == nil {
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

	if isEntryPointAddress(*destination) && functionSelector == functionSelectorEPHandleOps {
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

	if isEntryPointAddress(*destination) && hardForkTwoForbiddenEPFunctions(functionSelector) {
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

	gasLimit := uint64(100000)
	value := uint256.NewInt(0)

	ret, _, err := st.evm.Call(vm.AccountRef(from), appRegistryAddress, input, gasLimit, value)
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

	return result, nil
}

func isEntryPointAddress(address common.Address) bool {
	return address == aaEntryPointEnvAddress || address == aaEntryPointEnvAddressV7
}
