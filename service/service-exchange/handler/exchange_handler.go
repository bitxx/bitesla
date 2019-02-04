package handler

import (
	"context"
	"github.com/jason-wj/bitesla/service/service-exchange/conf"
	"github.com/jason-wj/bitesla/service/service-exchange/exchange/builder"
	"github.com/jason-wj/bitesla/service/service-exchange/proto"
)

var exchangeBuilder *builder.APIBuilder

type ExchangeHandler struct {
	repo *exchangeResposity
}

func NewExchangeHandler() *ExchangeHandler {
	exchangeBuilder = builder.NewAPIBuilder(conf.CurrentConfig.ServerConf.Proxy)
	repository := &exchangeResposity{}
	handler := &ExchangeHandler{
		repo: repository,
	}
	return handler
}

func (exchange *ExchangeHandler) LimitBuy(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	return exchange.repo.OrderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) LimitSell(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	return exchange.repo.OrderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) MarketBuy(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	return exchange.repo.OrderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) MarketSell(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	return exchange.repo.OrderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) CancelOrder(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, b *bitesla_srv_trader.Boolean) error {
	return exchange.repo.CancelOrder(reqCurrency, b)
}

func (exchange *ExchangeHandler) GetOneOrder(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	return exchange.repo.GetOneOrder(reqCurrency, order)
}

func (exchange *ExchangeHandler) GetUnfinishOrders(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Orders) error {
	return exchange.repo.GetUnfinishOrders(reqCurrency, order)
}

func (exchange *ExchangeHandler) GetOrderHistorys(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Orders) error {
	return exchange.repo.GetOrderHistorys(reqCurrency, order)
}

func (exchange *ExchangeHandler) GetTicker(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, tracker *bitesla_srv_trader.Ticker) error {
	return exchange.repo.GetTicker(reqCurrency, tracker)
}

func (exchange *ExchangeHandler) GetDepth(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, depth *bitesla_srv_trader.Depth) error {
	return exchange.repo.GetDepth(reqCurrency, depth)
}

func (exchange *ExchangeHandler) GetKlineRecords(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, respKlines *bitesla_srv_trader.Klines) error {
	return exchange.repo.getKlineRecords(reqCurrency, respKlines)
}

func (exchange *ExchangeHandler) GetTrades(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, trades *bitesla_srv_trader.Trades) error {
	return exchange.repo.GetTrades(reqCurrency, trades)
}

func (exchange *ExchangeHandler) GetAccount(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, account *bitesla_srv_trader.Accounts) error {
	return exchange.repo.getAccount(reqCurrency, account)
}

func (exchange *ExchangeHandler) GetExchangeName(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, name *bitesla_srv_trader.Str) error {
	return exchange.repo.GetExchangeName(reqCurrency, name)
}
