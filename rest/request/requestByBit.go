package request

import (
	bybitRest "github.com/TestingAccMar/CCXT_beYANG_ByBit/bybit/spot/v3/rest"
)

func BybitToWalletBalance(data interface{}) bybitRest.WalletBalance {
	bt, _ := data.(bybitRest.WalletBalance)
	return bt
}
