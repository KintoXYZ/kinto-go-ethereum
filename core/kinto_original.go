package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// Valid Kinto addresses before the hardfork
var originalKintoAddresses = map[common.Address]bool{
	aaEntryPointEnvAddress: true, // aaEntryPointEnvAddress
	kintoIdEnvAddress:      true, // kintoIdEnvAddress
	walletFactoryAddress:   true, // walletFactoryAddress
	paymasterAddress:       true, // paymasterAddress
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
