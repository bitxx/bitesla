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

func (exchange *ExchangeHandler) ListExchange(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, respCurrencys *bitesla_srv_exchange.Currencys) error {
	return exchange.repo.listExchange(reqCurrency, respCurrencys)
}

func (exchange *ExchangeHandler) PutExchange(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, respCurrency *bitesla_srv_exchange.Currency) error {
	return exchange.repo.putExchange(reqCurrency, respCurrency)
}

func (exchange *ExchangeHandler) DeleteExchange(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, respCurrency *bitesla_srv_exchange.Currency) error {
	return exchange.repo.deleteExchange(reqCurrency, respCurrency)
}

func (exchange *ExchangeHandler) LimitBuy(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	return exchange.repo.orderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) LimitSell(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	return exchange.repo.orderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) MarketBuy(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	return exchange.repo.orderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) MarketSell(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	return exchange.repo.orderPlace(reqCurrency, order)
}

func (exchange *ExchangeHandler) CancelOrder(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, b *bitesla_srv_exchange.Boolean) error {
	return exchange.repo.cancelOrder(reqCurrency, b)
}

func (exchange *ExchangeHandler) GetOneOrder(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	return exchange.repo.getOneOrder(reqCurrency, order)
}

func (exchange *ExchangeHandler) GetUnfinishOrders(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Orders) error {
	return exchange.repo.getUnfinishOrders(reqCurrency, order)
}

func (exchange *ExchangeHandler) GetOrderHistorys(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Orders) error {
	return exchange.repo.getOrderHistorys(reqCurrency, order)
}

func (exchange *ExchangeHandler) GetTicker(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, tracker *bitesla_srv_exchange.Ticker) error {
	return exchange.repo.getTicker(reqCurrency, tracker)
}

func (exchange *ExchangeHandler) GetDepth(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, depth *bitesla_srv_exchange.Depth) error {
	err := exchange.repo.getDepth(reqCurrency, depth)
	return err
}

func (exchange *ExchangeHandler) GetKlineRecords(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, respKlines *bitesla_srv_exchange.Klines) error {
	return exchange.repo.getKlineRecords(reqCurrency, respKlines)
}

func (exchange *ExchangeHandler) GetTrades(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, trades *bitesla_srv_exchange.Trades) error {
	return exchange.repo.getTrades(reqCurrency, trades)
}

func (exchange *ExchangeHandler) GetAccount(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, account *bitesla_srv_exchange.Accounts) error {
	return exchange.repo.getAccount(reqCurrency, account)
}

func (exchange *ExchangeHandler) GetExchangeDetail(ctx context.Context, reqCurrency *bitesla_srv_exchange.Currency, respCurrency *bitesla_srv_exchange.Currency) error {
	return exchange.repo.getExchangeDetail(reqCurrency, respCurrency)
}
