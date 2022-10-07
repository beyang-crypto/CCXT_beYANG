package request

import (
	ftxRest "github.com/TestingAccMar/CCXT_beYANG_FTX/ftx/spot/rest"
)

func FTXToWalletBalance(data interface{}) ftxRest.WalletBalance {
	bt, _ := data.(ftxRest.WalletBalance)
	return bt
}
