package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

var hardfork4KintoAddresses = map[common.Address]bool{
	aaEntryPointEnvAddress:      true, // aaEntryPointEnvAddress
	kintoIdEnvAddress:           true, // kintoIdEnvAddress
	walletFactoryAddress:        true, // walletFactoryAddress
	paymasterAddress:            true, // paymasterAddress
	appRegistryAddress:          true, // appRegistryAddress
	upgradeExecutor:             true, // upgradeExecutor
	customGatewayAddress:        true, // customGatewayAddress
	gatewayRouterAddress:        true, // gatewayRouterAddress
	standardGatewayAddress:      true, // standardGatewayAddress
	wethGateWayAddress:          true, // wethGateWayAddress
	bundleBulker:                true, // bundleBulker
	arbRetrayableTx:             true, // arbRetrayableTx
	socket:                      true, // socket
	socketExecutionManager:      true, // socketExecutionManager
	socketTransmitManager:       true, // socketTransmitManager
	socketFastSwitchboard:       true, // socketFastSwitchboard
	socketOptimisticSwitchboard: true, // socketOptimisticSwitchboard
	socketBatcher:               true, // socketBatcher
	socketSimulator:             true, // socketSimulator
	socketSimulatorUtils:        true, // socketSimulatorUtils
	socketSwitchboardSimulator:  true, // socketSwitchboardSimulator
	socketCapacitorSimulator:    true, // socketCapacitorSimulator
	create2Factory:              true, // create2Factory
}

func enforceHardForkFourRules(msg *Message) error {
	if msg.TxRunMode == MessageGasEstimationMode {
		return nil // allow gas estimation
	}

	destination := msg.To
	functionSelector := extractFunctionSelector(msg.Data)

	if destination == nil {
		return fmt.Errorf("%w: %v EOAs can't create contracts directly, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if _, ok := hardfork4KintoAddresses[*destination]; !ok {
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
