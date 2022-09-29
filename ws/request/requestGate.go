package request

import (
	gate "github.com/TestingAccMar/CCXT_beYANG_Gate/gate/spotAndMargin/v4/ws"
)

func GateTickers(data interface{}) gate.Tickers {
	t, _ := data.(gate.Tickers)
	return t
}
