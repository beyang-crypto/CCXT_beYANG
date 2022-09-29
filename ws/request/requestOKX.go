package request

import okxWs "github.com/TestingAccMar/CCXT_beYANG_OKX/okx/ws"

func OKXTickers(data interface{}) okxWs.Tickers {
	t, _ := data.(okxWs.Tickers)
	return t
}
