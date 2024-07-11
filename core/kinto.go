package core

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

// Kinto addresses mainnet
/*
var (
	aaEntryPointEnvAddress      = common.HexToAddress("0x2843C269D2a64eCfA63548E8B3Fc0FD23B7F70cb")
	kintoIdEnvAddress           = common.HexToAddress("0xf369f78E3A0492CC4e96a90dae0728A38498e9c7")
	walletFactoryAddress        = common.HexToAddress("0x8a4720488CA32f1223ccFE5A087e250fE3BC5D75")
	paymasterAddress            = common.HexToAddress("0x1842a4EFf3eFd24c50B63c3CF89cECEe245Fc2bd")
	appRegistryAddress          = common.HexToAddress("0x5A2b641b84b0230C8e75F55d5afd27f4Dbd59d5b")
	upgradeExecutor             = common.HexToAddress("0x88e03D41a6EAA9A0B93B0e2d6F1B34619cC4319b")
	customGatewayAddress        = common.HexToAddress("0x06FcD8264caF5c28D86eb4630c20004aa1faAaA8")
	gatewayRouterAddress        = common.HexToAddress("0x340487b92808B84c2bd97C87B590EE81267E04a7")
	standardGatewayAddress      = common.HexToAddress("0x87799989341A07F495287B1433eea98398FD73aA")
	wethGateWayAddress          = common.HexToAddress("0xd563ECBDF90EBA783d0a218EFf158C1263ad02BE")
	bundleBulker                = common.HexToAddress("0x8d2D899402ed84b6c0510bB1ad34ee436ADDD20d")
	arbRetrayableTx             = common.HexToAddress("0x000000000000000000000000000000000000006E")
	socket                      = common.HexToAddress("0x3e9727470C66B1e77034590926CDe0242B5A3dCc")
	socketExecutionManager      = common.HexToAddress("0x6c914cc610e9a05eaFFfD79c10c60Ad1704717E5")
	socketTransmitManager       = common.HexToAddress("0x6332e56A423480A211E301Cb85be12814e9238Bb")
	socketFastSwitchboard       = common.HexToAddress("0x516302D1b25e5F6d1ac90eF7256270cd799524CF")
	socketOptimisticSwitchboard = common.HexToAddress("0x2B98775aBE9cDEb041e3c2E56C76ce2560AF57FB")
	socketBatcher               = common.HexToAddress("0x12FF8947a2524303C13ca7dA9bE4914381f6557a")
	socketSimulator             = common.HexToAddress("0x72846179EF1467B2b71F2bb7525fcD4450E46B2A")
	socketSimulatorUtils        = common.HexToAddress("0x897DA4D039f64090bfdb33cd2Ed2Da81adD6FB02")
	socketSwitchboardSimulator  = common.HexToAddress("0xa7527C270f30cF3dAFa6e82603b4978e1A849359")
	socketCapacitorSimulator    = common.HexToAddress("0x6dbB5ee7c63775013FaF810527DBeDe2810d7Aee")
	create2Factory              = common.HexToAddress("0x4e59b44847b379578588920cA78FbF26c0B4956C")
	aaEntryPointEnvAddressV7    = common.HexToAddress("0x0000000071727De22E5E9d8BAf0edAc6f37da032")
)
*/
// Kinto addresses devnet

var (
	aaEntryPointEnvAddress      = common.HexToAddress("0x691aC5BA3cb64CF5b8d4a6484f933794E2dF5d40")
	kintoIdEnvAddress           = common.HexToAddress("0xCF71C996cD870069Aba049525a445c5B79020a53")
	walletFactoryAddress        = common.HexToAddress("0x537fA09ef76BB964D0C9dfDdff5552706DfadbC0")
	paymasterAddress            = common.HexToAddress("0x0dc36bac72A99d70Fa8f2CB3f780e511a691841b")
	appRegistryAddress          = common.HexToAddress("0xC9524e5C6Bd274fEb8cea7BaB7e3Ac7b06F5a190")
	upgradeExecutor             = common.HexToAddress("0x6B0d3F40DeD9720938DB274f752F1e11532c2640")
	customGatewayAddress        = common.HexToAddress("0x094F8C3eA1b5671dd19E15eCD93C80d2A33fCA99")
	gatewayRouterAddress        = common.HexToAddress("0xf3AC740Fcc64eEd76dFaE663807749189A332d54")
	standardGatewayAddress      = common.HexToAddress("0x6A8d32c495df943212B7788114e41103047150a5")
	wethGateWayAddress          = common.HexToAddress("0x79B47F0695608aD8dc90E400a3E123b02eB72D24")
	bundleBulker                = common.HexToAddress("0x2291d967F4f8E7B062D0eAA977C5adBbd33B99BB")
	arbRetrayableTx             = common.HexToAddress("0x000000000000000000000000000000000000006E")
	socket                      = common.HexToAddress("0x62B421B7dbc6207CC010318a4ba567786137de29")
	socketExecutionManager      = common.HexToAddress("0x4518D09052D6f40f83d489a3E9F81EF369dB0753")
	socketTransmitManager       = common.HexToAddress("0x956b0c4d2f3f050bDB6A5b6B6a95050af9fA3A62")
	socketFastSwitchboard       = common.HexToAddress("0x38cACa8a8b5579Cb2d2870A73DbfAa54B6Ee490D")
	socketOptimisticSwitchboard = common.HexToAddress("0xC94De3804d3c67620E7a70547bCB4a77b53952EC")
	socketBatcher               = common.HexToAddress("0x1d1ef33231689d6057565f99d8B1864E6bE5eb94")
	socketSimulator             = common.HexToAddress("0xbfd616DA87ebea4513aB633C9298218dd4a698dc")
	socketSimulatorUtils        = common.HexToAddress("0x93f73A15272D4D46720234C32BC1eE7290Eb5F18")
	socketSwitchboardSimulator  = common.HexToAddress("0x108eE40304fB1C3560eFF91f8E15B52ea4E2a257")
	socketCapacitorSimulator    = common.HexToAddress("0x1390e33B8F1D6D92e27fcEF2c6E5641Be951A2bb")
	create2Factory              = common.HexToAddress("0x4e59b44847b379578588920cA78FbF26c0B4956C")
	aaEntryPointEnvAddressV7    = common.HexToAddress("0x0000000071727De22E5E9d8BAf0edAc6f37da032")
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

// enforceKinto decides which set of Kinto rules to apply based on the current block number
func enforceKinto(msg *Message, st *StateTransition) error {
	var currentBlockNumber = st.evm.Context.BlockNumber

	// Hardfork5 bytecode replacement (happens once)
	if currentBlockNumber.Cmp(common.KintoHardfork5) == 0 {
		st.state.SetCode(replaceHF5Address, replacedHF5Bytecode)
	}

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
		} else if currentBlockNumber.Cmp(common.KintoHardfork4) <= 0 {
			return enforceHardForkThreeRules(msg) // Rules for the third hard fork
		} else if currentBlockNumber.Cmp(common.KintoHardfork5) <= 0 {
			return enforceHardForkFourRules(msg) // Rules for the fourth hard fork
		} else if currentBlockNumber.Cmp(common.KintoHardfork6) <= 0 {
			return enforceHardForkFiveRules(msg) // Rules for the fifth hard fork
		} else {
			return enforceHardForkSixRules(st) // Rules for the sixth hard fork
		}
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
