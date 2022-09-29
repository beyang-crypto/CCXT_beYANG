package request

import (
	okxRest "github.com/TestingAccMar/CCXT_beYANG_OKX/okx/rest"
)

func OKXToWalletBalance(data interface{}) okxRest.WalletBalance {
	bt, _ := data.(okxRest.WalletBalance)
	return bt
}
