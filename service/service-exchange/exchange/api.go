package exchange

import "github.com/jason-wj/bitesla/service/service-exchange/proto"

type Api interface {
	LimitBuy(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Order) error
	LimitSell(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Order) error
	MarketBuy(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Order) error
	MarketSell(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Order) error
	CancelOrder(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Boolean) error
	GetOneOrder(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Order) error
	GetUnfinishOrders(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Orders) error
	GetOrderHistorys(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Orders) error
	GetTicker(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Ticker) error
	GetDepth(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Depth) error
	GetKlineRecords(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Klines) error
	GetTrades(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Trades) error
	GetAccount(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Accounts) error
}
