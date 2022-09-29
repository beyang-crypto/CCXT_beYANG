package request

import (
	ftxWs "github.com/TestingAccMar/CCXT_beYANG_FTX/ftx/ws"
)

func FTXTicker(data interface{}) ftxWs.Ticker {
	t, _ := data.(ftxWs.Ticker)
	return t
}
