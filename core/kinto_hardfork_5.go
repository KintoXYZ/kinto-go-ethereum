package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

var replaceHF5Address = common.HexToAddress("0x4e59b44847b379578588920cA78FbF26c0B4956C")
var replacedHF5Bytecode = common.FromHex("0x6080604052600436106100225760003560e01c806399a6cddd146101225761004d565b3661004d57604051638dc2b20960e01b81523360048201523460248201526044015b60405180910390fd5b6040516313289ea360e31b815233600482015260009036906060906001600160a01b037f000000000000000000000000f369f78e3a0492cc4e96a90dae0728a38498e9c71690639944f51890602401602060405180830381865afa1580156100b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100dd9190610172565b6100fc57604051630ab529c360e21b8152336004820152602401610044565b36601f1901600081602082378035828234f5915081610119578081fd5b8181526014600cf35b34801561012e57600080fd5b506101567f000000000000000000000000f369f78e3a0492cc4e96a90dae0728a38498e9c781565b6040516001600160a01b03909116815260200160405180910390f35b60006020828403121561018457600080fd5b8151801515811461019457600080fd5b939250505056fea2646970667358221220a35916f465fc2e7b19efa5b0d984bc014269d1d8cbad39f684ed4ecc1dd45e5f64736f6c63430008180033")

var hardfork5KintoAddresses = map[common.Address]bool{
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

var (
	dinaryStagingEOAs = []common.Address{
		//owner
		common.HexToAddress("0x4181803232280371E02a875F51515BE57B215231"),
		//Kinto d-shares operator/keepers
		common.HexToAddress("0x8D69Ec4029c7d634a06FaB50da4b538499FC8598"), //also in usdplus
		common.HexToAddress("0x874c1606c678cdA1d0f054f5123567198B13fedF"), //also in usdplus
		common.HexToAddress("0xECC40Cf598B1e98846267F274559062aE4cd3F9D"),
		common.HexToAddress("0xd9ADdFcb54cC09902C49E02BdD9Ad05a003dA630"),

		//universal
		common.HexToAddress("0xf4ce0c560dEbcA25F1daA0B082Ffb6B3E3B66B3C"), //minter
	}

	dinaryStagingContracts = []common.Address{
		//Kinto d-shares
		common.HexToAddress("0xF34f9C994E28254334C83AcE353d814E5fB90815"),
		common.HexToAddress("0x17C477f860aD70541277eF59D5c55aaB0137dbB8"),
		common.HexToAddress("0x2e92D8Ba4122a40922BE2B46E01982749d8FC127"),
		common.HexToAddress("0x5fc67f2EE4e30D020A930B745aaDb68DDa985a4C"),
		common.HexToAddress("0xB621dA3AFC9Df83209042De965dD4Ccb0e8a0ABA"),
		common.HexToAddress("0x251b1B7c4957FB9Db75921E50F4cf2a5e284b224"),
		common.HexToAddress("0xA4DbdcEFFCbc6141C88F08b3D455775B34218250"),
		common.HexToAddress("0xdA25A48456bBdbBe41a03B0D50ba74993A8A0Fa0"),

		//usdplus
		common.HexToAddress("0x7031b2EA8B97304885b8c842E14BFc5DD6FC92f8"),
		common.HexToAddress("0x0a511eC63c836037F0A2CcC0A81984247E27783b"),
		common.HexToAddress("0xa7D259925f951b674bCDbcF7a63Ab2f5923483dB"),
		common.HexToAddress("0x2eeBEa5eb4a0feA2ec20FD48A2289D87E2882C71"),
		common.HexToAddress("0x90AB5E52Dfcce749CA062f4e04292fd8a67E86b3"),

		//universal
		common.HexToAddress("0x09E365aCDB0d936DD250351aD0E7de3Dad8706E5"),
		common.HexToAddress("0xC60bB79d0176d9C2FD23Eaeff91AC800b3ae5A83"),
	}

	dinaryProductionEOAs = []common.Address{
		//owner
		common.HexToAddress("0x269e944aD9140fc6e21794e8eA71cE1AfBfe38c8"),

		//Kinto d-shares operator/keepers
		common.HexToAddress("0x2bF22fD411C71b698bF6e0e937b1B948339Ec369"), //also in usdplus
		common.HexToAddress("0x0556Fe4ddffE2798BA34E1A92306B12cBC6c94fC"), //also in usdplus
		common.HexToAddress("0xAa0ed80DE46CF02bde4493A84FE22Af8fE79c01f"),
		common.HexToAddress("0x0D5e0d9717998059cB34945dC231f7619107E53e"),

		//universal
		common.HexToAddress("0x334d41e2705D8a41999A48Fddd2a06F146C4AF59"), //minter

	}

	dinaryProductionContracts = []common.Address{
		//Kinto d-shares
		common.HexToAddress("0xB2eEc63Cdc175d6d07B8f69804C0Ab5F66aCC3cb"),
		common.HexToAddress("0xa9a60Ccc6363e440eeEaa8Ad015607c7a34360CE"),
		common.HexToAddress("0xd1d93E6Ad5219083Bb2cf3B065a562223381b71F"),
		common.HexToAddress("0xE4Daa69e99F48AD0C4D4843deF4447253248A906"),
		common.HexToAddress("0x1498A49Ff90d9f7fE8915658A1FC3b87c9A4Ba8c"),
		common.HexToAddress("0xa089dC07A4baFd941a4323a9078D2c24be8A747C"),
		common.HexToAddress("0x1464727DCC5619E430FaA217a61180d1cEDd2d3a"),
		common.HexToAddress("0x8E58548731Ae14D573b54647f2dc393639519fF3"),

		//usdplus
		common.HexToAddress("0xd4ee24378201190c7C50D52D3D29C459a1278F91"),
		common.HexToAddress("0x6F086dB0f6A621a915bC90295175065c9e5d9b8c"),
		common.HexToAddress("0xeDA274898ED364Bd346fA74cf6eCAB4BF8f1665f"),
		common.HexToAddress("0x931C5dC9eA13b0F6B4768a98AFfEA773b888e978"),

		//Universal
		common.HexToAddress("0x400880b800410B2951Afd0503dC457aea8A4bAb5"),
		common.HexToAddress("0xF96b974FE330C29e80121E33ed4071C283257979"),
		common.HexToAddress("0xAdFeB630a6aaFf7161E200088B02Cf41112f8B98"),
	}
)

func enforceHardForkFiveRules(msg *Message) error {
	if msg.TxRunMode == MessageGasEstimationMode {
		return nil // allow gas estimation
	}

	destination := msg.To
	origin := msg.From

	if (containsAddress(dinaryStagingEOAs, origin) && containsAddress(dinaryStagingContracts, *destination)) ||
		(containsAddress(dinaryProductionEOAs, origin) && containsAddress(dinaryProductionContracts, *destination)) {
		return nil // allow dinary contracts to interact with dinary EOAs
	}

	functionSelector := extractFunctionSelector(msg.Data)

	if destination == nil {
		return fmt.Errorf("%w: %v EOAs can't create contracts directly, %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	if _, ok := hardfork5KintoAddresses[*destination]; !ok {
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

func containsAddress(slice []common.Address, addr common.Address) bool {
	for _, a := range slice {
		if a == addr {
			return true
		}
	}
	return false
}
