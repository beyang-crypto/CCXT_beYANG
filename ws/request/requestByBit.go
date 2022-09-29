package request

import (
	bybit "github.com/TestingAccMar/CCXT_beYANG_ByBit/bybit/spot/v3/ws"
)

func BybitBookTicker(data interface{}) bybit.BookTicker {
	bt, _ := data.(bybit.BookTicker)
	return bt
}
