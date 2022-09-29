package main

import (
	"time"

	"github.com/TestingAccMar/CCXT_beYANG/rest"
	restconsts "github.com/TestingAccMar/CCXT_beYANG/rest/consts"
	restrequest "github.com/TestingAccMar/CCXT_beYANG/rest/request"

	"log"
)

func main() {

	exResOKX := rest.ExchangeRest{
		Name:           "OKX",
		Addr:           restconsts.OKXRestURL,
		ApiKey:         "da092be5-da8f-4268-a786-f2dbd296dc27",
		SecretKey:      "3602F5C69EBCC49BE99C4F31EB1D11B7",
		APIKeyPassword: "Pravdavsile/1",
		DebugMode:      true,
	}

	okxRestEx := rest.NewExchange(exResOKX)
	go func() {
		time.Sleep(1 * time.Second)
		balance := restrequest.OKXToWalletBalance(okxRestEx.GetBalance())
		for _, datas := range balance.Data {
			for _, details := range datas.Details {
				log.Printf("coin = %s, total = %s", details.Ccy, details.CashBal)
			}
		}
	}()

	forever := make(chan struct{})
	<-forever
}
