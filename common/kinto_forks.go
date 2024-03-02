package common

import (
	"math/big"
)

// Block numbers for Kinto rule changes
var (
	KintoRulesBlockStart = big.NewInt(100)
	KintoHardfork1       = big.NewInt(150)
	KintoHardfork2       = big.NewInt(200)
)