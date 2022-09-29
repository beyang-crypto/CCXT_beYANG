package consts

import okxWs "github.com/TestingAccMar/CCXT_beYANG_OKX/okx/ws"

const (
	OKXHostPublicWebSocket     = okxWs.HostPublicWebSocket
	OKXHostPrivateWebSocket    = okxWs.HostPrivateWebSocket
	OKXHostPublicWebSocketAWS  = okxWs.HostPublicWebSocketAWS
	OKXHostPrivateWebSocketAWS = okxWs.HostPrivateWebSocketAWS
)

const (
	OKXChannelBalanceAndPosition = okxWs.ChannelBalanceAndPosition
	OKXChannelRfqs               = okxWs.ChannelRfqs
	OKXChannelQuotes             = okxWs.ChannelQuotes
	OKXChannelTicker             = okxWs.ChannelTicker
	OKXChannelOptSummary         = okxWs.ChannelOptSummary
	OKXChannelAccount            = okxWs.ChannelAccount
	OKXChannelAccountGreeks      = okxWs.ChannelAccountGreeks
	OKXChannelGridPositions      = okxWs.ChannelGridPositions
	OKXChannelGridSubOrders      = okxWs.ChannelGridSubOrders
	OKXChannelLiquidationWarning = okxWs.ChannelLiquidationWarning
	OKXChannelGridOrdersSpot     = okxWs.ChannelGridOrdersSpot
	OKXChannelGridOrdersContract = okxWs.ChannelGridOrdersContract
	OKXChannelGridOrdersMoon     = okxWs.ChannelGridOrdersMoon
	OKXChannelAlgoAdvance        = okxWs.ChannelAlgoAdvance
	OKXChannelEstimatedPrice     = okxWs.ChannelEstimatedPrice
	OKXChannelPositions          = okxWs.ChannelPositions
	OKXChannelOrders             = okxWs.ChannelOrders
	OKXChannelOrdersAlgo         = okxWs.ChannelOrdersAlgo
)

const (
	OKXInstTypeSpot    = okxWs.InstTypeSpot
	OKXInstTypeMargin  = okxWs.InstTypeMargin
	OKXInstTypeSwap    = okxWs.InstTypeSwap
	OKXInstTypeFutures = okxWs.InstTypeFutures
	OKXInstTypeOption  = okxWs.InstTypeOption
	OKXInstTypeAny     = okxWs.InstTypeAny
)
