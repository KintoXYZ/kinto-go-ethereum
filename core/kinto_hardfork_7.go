package core

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

func enforceHardForkSevenRules(st *StateTransition) error {
	msg := st.msg

	if msg.TxRunMode == MessageGasEstimationMode {
		return nil // allow gas estimation
	}

	destination := msg.To
	origin := msg.From

	if destination == nil {
		destination = &ZeroAddress
	}

	allowed, err := isContractCallAllowedFromEOAHF7(st, origin, *destination, st.msg.Data, st.msg.Value)

	if allowed && err == nil {
		return nil
	}

	if !allowed && err == nil {
		return fmt.Errorf("%w: %v is not allowed to call %v", ErrKintoNotAllowed, msg.From.Hex(), destination)
	}

	//if it is !allowed and err !=nil something went wrong with the contract call
	//however we still allow the transaction to proceed or it will brick the chain
	return nil
}

func isContractCallAllowedFromEOAHF7(st *StateTransition, from, to common.Address, data []byte, value *big.Int) (bool, error) {
	// Define the updated ABI
	const abiJSON = `[{"inputs":[{"internalType":"address","name":"sender","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"bytes","name":"callData","type":"bytes"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"isContractCallAllowedFromEOA","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return false, fmt.Errorf("error parsing ABI: %v", err)
	}

	// Pack the function call with the new parameters
	input, err := parsedABI.Pack("isContractCallAllowedFromEOA", from, to, data, value)
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

	return result, nil
}
