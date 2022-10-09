package main

import (
	"github.com/TestingAccMar/CCXT_beYANG/ws"
	wsconsts "github.com/TestingAccMar/CCXT_beYANG/ws/consts"
	wsrequest "github.com/TestingAccMar/CCXT_beYANG/ws/request"

	"log"
)

func main() {
	exBinance := ws.ExchangeWS{
		Name:      "binance",
		Addr:      wsconsts.BinanceHostMainnetPublicTopics,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: false,
	}

	exBybit := ws.ExchangeWS{
		Name:      "bybit",
		Addr:      wsconsts.BybitHostMainnetPublicTopics,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: false,
	}

	exFTX := ws.ExchangeWS{
		Name:      "ftx",
		Addr:      wsconsts.FTXHostMainnetPublicTopics,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: false,
	}

	exOKX := ws.ExchangeWS{
		Name:      "okx",
		Addr:      wsconsts.OKXHostPublicWebSocket,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: false,
	}
	exGate := ws.ExchangeWS{
		Name:      "gate",
		Addr:      wsconsts.GateHostWebsocketURL,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: false,
	}

	binanceEx := ws.NewExchange(exBinance)
	bybitEx := ws.NewExchange(exBybit)
	ftxEx := ws.NewExchange(exFTX)
	okxEx := ws.NewExchange(exOKX)
	gateEx := ws.NewExchange(exGate)

	ws.Start(binanceEx)
	ws.Start(bybitEx)
	ws.Start(ftxEx)
	ws.Start(okxEx)
	ws.Start(gateEx)

	pairBinance := ws.GetPair(binanceEx, "BTC", "USDT")
	pairBybit := ws.GetPair(bybitEx, "BTC", "USDT")
	pairFTX := ws.GetPair(ftxEx, "BTC", "USDT")
	pairOKX := ws.GetPair(okxEx, "BTC", "USDT")
	pairGate := ws.GetPair(gateEx, "BTC", "USDT")

	ws.Subscribe(binanceEx, wsconsts.BinanceChannelTicker, []string{pairBinance})
	ws.Subscribe(bybitEx, wsconsts.BybitChannelBookTicker, []string{pairBybit})
	ws.Subscribe(ftxEx, wsconsts.FTXChannelTicker, []string{pairFTX})
	ws.Subscribe(okxEx, wsconsts.OKXChannelTicker, []string{pairOKX})
	ws.Subscribe(gateEx, wsconsts.GateChannelTicker, []string{pairGate})

	ws.On(binanceEx, wsconsts.BinanceChannelTicker, handleBookTicker)
	ws.On(bybitEx, wsconsts.BybitChannelBookTicker, handleBookTicker)
	ws.On(ftxEx, wsconsts.FTXChannelTicker, handleBookTicker)
	ws.On(okxEx, wsconsts.OKXChannelTicker, handleBookTicker)
	ws.On(gateEx, wsconsts.GateChannelTicker, handleBookTicker)

	forever := make(chan struct{})
	<-forever
}

func handleBookTicker(name string, symbol string, data interface{}) {
	switch name {
	case "Binance":
		{
			ticker, ok := wsrequest.BinanceToBookTicker(data)
			if ok {
				log.Printf("%s Ticker  %s: %v", name, symbol, ticker.B)
			}
		}
	case "Bybit":
		{
			ticker, ok := wsrequest.BybitBookTicker(data)
			if ok {
				log.Printf("%s Ticker  %s: %v", name, symbol, ticker.Data.Bp)
			}
		}
	case "FTX":
		{
			ticker, ok := wsrequest.FTXTicker(data)
			if ok {
				log.Printf("%s Ticker  %s: %v", name, symbol, ticker.Data.Bid)
			}
		}
	case "OKX":
		{
			ticker, ok := wsrequest.OKXTickers(data)
			if ok {
				log.Printf("%s Ticker  %s: %v", name, symbol, ticker.Data[0].BidPx)
			}
		}

	case "Gate":
		{
			ticker, ok := wsrequest.GateTickers(data)
			if ok {
				log.Printf("%s Ticker  %s: %v", name, symbol, ticker.Result.HighestBid)
			}
		}
	}

}
