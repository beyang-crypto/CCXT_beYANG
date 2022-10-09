package request

import okxWs "github.com/TestingAccMar/CCXT_beYANG_OKX/okx/ws"

func OKXTickers(data interface{}) (okxWs.Tickers, bool) {
	t, ok := data.(okxWs.Tickers)
	return t, ok
}
