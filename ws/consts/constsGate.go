package consts

import (
	gate "github.com/TestingAccMar/CCXT_beYANG_Gate/gate/spotAndMargin/v4/ws"
)

const (
	GateHostWebsocketURL = gate.HostWebsocketURL
)

const (
	//https://www.gate.io/docs/developers/apiv4/ws/en/#tickers-channel
	GateChannelTicker = gate.ChannelTicker

	// https://www.gate.io/docs/developers/apiv4/ws/en/#spot-balance-channel
	GateChannelBalances = gate.ChannelBalances
)
