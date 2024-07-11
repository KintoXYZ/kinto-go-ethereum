package common

import (
	"math/big"
)

// Block numbers for Kinto rule changes
var (
	KintoRulesBlockStart = big.NewInt(100) //100
	KintoHardfork1       = big.NewInt(120) //57000
	KintoHardfork2       = big.NewInt(140) //118000
	KintoHardfork3       = big.NewInt(160) //125000
	KintoHardfork4       = big.NewInt(180) //133000
	KintoHardfork5       = big.NewInt(200)
	KintoHardfork6       = big.NewInt(220)
	SelfDestructWallet   = "0x660ad4B5A74130a4796B4d54BC6750Ae93C86e6c"
)
