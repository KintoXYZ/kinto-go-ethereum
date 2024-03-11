package common

import (
	"math/big"
)

// Block numbers for Kinto rule changes
var (
	KintoRulesBlockStart = big.NewInt(100)
	KintoHardfork1       = big.NewInt(57000)
	KintoHardfork2       = big.NewInt(118000)
	SelfDestructWallet   = "0x660ad4B5A74130a4796B4d54BC6750Ae93C86e6c"
)