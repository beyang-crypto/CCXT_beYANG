package request

import (
	gate "github.com/TestingAccMar/CCXT_beYANG_Gate/gate/spotAndMargin/v4/rest"
)

func GateToWalletDeposits(data interface{}) gate.WalletDeposits {
	bt, _ := data.(gate.WalletDeposits)
	return bt
}
