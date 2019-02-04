package exchange

import "github.com/jason-wj/bitesla/service/service-exchange/proto"

// api interface

type API interface {
	LimitBuy(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error)
	LimitSell(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error)
	MarketBuy(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error)
	MarketSell(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error)
	CancelOrder(orderId string, currency string) (bool, error)
	GetOneOrder(orderId string, currency string) (*bitesla_srv_trader.Order, error)
	GetUnfinishOrders(currency string) ([]*bitesla_srv_trader.Order, error)
	GetOrderHistorys(currency string, currentPage, pageSize int32) ([]*bitesla_srv_trader.Order, error)
	GetAccount() ([]*bitesla_srv_trader.Account, error)

	GetTicker(currency string) (*bitesla_srv_trader.Ticker, error)
	GetDepth(size int32, currency string) (*bitesla_srv_trader.Depth, error)
	GetKlineRecords(currency string, period, size, since int32) ([]*bitesla_srv_trader.Kline, error)
	//非个人，整个交易所的交易记录
	GetTrades(currencyPair string, since int32) ([]*bitesla_srv_trader.Trade, error)

	GetExchangeName() string
}
