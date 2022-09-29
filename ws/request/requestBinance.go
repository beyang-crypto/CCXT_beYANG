package request

import (
	binanceWs "github.com/TestingAccMar/CCXT_beYANG_Binance/binance/ws"
)

func BinanceToBookTicker(data interface{}) binanceWs.BookTicker {
	bt, _ := data.(binanceWs.BookTicker)
	return bt
}
