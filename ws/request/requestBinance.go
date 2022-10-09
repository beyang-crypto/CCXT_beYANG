package request

import (
	binanceWs "github.com/TestingAccMar/CCXT_beYANG_Binance/binance/ws"
)

func BinanceToBookTicker(data interface{}) (binanceWs.BookTicker, bool) {
	bt, ok := data.(binanceWs.BookTicker)
	return bt, ok
}
