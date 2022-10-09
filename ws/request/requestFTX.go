package request

import (
	ftxWs "github.com/TestingAccMar/CCXT_beYANG_FTX/ftx/spot/ws"
)

func FTXTicker(data interface{}) (ftxWs.Ticker, bool) {
	t, ok := data.(ftxWs.Ticker)
	return t, ok
}
