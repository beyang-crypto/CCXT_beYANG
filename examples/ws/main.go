package main

import (
	"github.com/TestingAccMar/CCXT_beYANG/ws"
	wsconsts "github.com/TestingAccMar/CCXT_beYANG/ws/consts"
	wsrequest "github.com/TestingAccMar/CCXT_beYANG/ws/request"

	"log"
)

func main() {
	exFTX := ws.ExchangeWS{
		Name:      "binance",
		Addr:      wsconsts.BinanceHostMainnetPublicTopics,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: false,
	}

	binanceEx := ws.NewExchange(exFTX)
	ws.Start(binanceEx)

	pairOKX := ws.GetPair(binanceEx, "BTC", "USDT")
	ws.Subscribe(binanceEx, wsconsts.BinanceChannelTicker, []string{pairOKX})

	ws.On(binanceEx, wsconsts.BinanceChannelTicker, handleBookTickerBinace)

	forever := make(chan struct{})
	<-forever
}

func handleBookTickerOKX(symbol string, data interface{}) {
	aa, ok := wsrequest.OKXTickers(data)
	if ok {
		log.Printf("OKX Ticker  %s: %v", symbol, aa.Data[0].BidPx)
	}
}

func handleBookTickerBinace(symbol string, data interface{}) {
	aa, ok := wsrequest.BinanceToBookTicker(data)
	if ok {
		log.Printf("Binance Ticker  %s: %v", symbol, aa.B)
	}
}
