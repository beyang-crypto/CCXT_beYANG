package main

import (
	"github.com/TestingAccMar/CCXT_beYANG/ws"
	wsconsts "github.com/TestingAccMar/CCXT_beYANG/ws/consts"
	wsrequest "github.com/TestingAccMar/CCXT_beYANG/ws/request"

	"log"
)

func main() {
	exFTX := ws.ExchangeWS{
		Name:      "OKX",
		Addr:      wsconsts.OKXHostPublicWebSocket,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: true,
	}

	okxEx := ws.NewExchange(exFTX)
	ws.Start(okxEx)

	pairOKX := ws.GetPair(okxEx, "BTC", "USDT")
	ws.Subscribe(okxEx, wsconsts.OKXChannelTicker, pairOKX)

	ws.On(okxEx, wsconsts.OKXChannelTicker, handleBookTickerOKX)

	forever := make(chan struct{})
	<-forever
}

func handleBookTickerOKX(symbol string, data interface{}) {
	aa := wsrequest.OKXTickers(data)
	log.Printf("OKX Ticker  %s: %v", symbol, aa.Data[0].BidPx)
}
