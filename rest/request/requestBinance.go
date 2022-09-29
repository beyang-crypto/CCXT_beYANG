package request

import (
	binanceRest "github.com/TestingAccMar/CCXT_beYANG_Binance/binance/rest"
)

func BinanceToWalletBalance(data interface{}) binanceRest.WalletBalance {
	bt, _ := data.(binanceRest.WalletBalance)
	return bt
}
