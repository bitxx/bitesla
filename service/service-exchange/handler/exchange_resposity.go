package handler

import (
	"context"
	"errors"
	"github.com/jason-wj/bitesla/service/service-exchange/proto"
)

type exchangeResposity struct {
}

func (e *exchangeResposity) getKlineRecords(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, kLines *bitesla_srv_trader.Klines) (err error) {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	return api.GetKlineRecords(ctx, reqCurrency, kLines)
}

func (e *exchangeResposity) getAccount(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, account *bitesla_srv_trader.Accounts) (err error) {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	return api.GetAccount(ctx, reqCurrency, account)
}

func (e *exchangeResposity) OrderPlace(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}

	switch reqCurrency.OrderType {
	case int32(bitesla_srv_trader.TradeSide_BUY):
		err = api.LimitBuy(ctx, reqCurrency, order)
	case int32(bitesla_srv_trader.TradeSide_SELL):
		err = api.LimitSell(ctx, reqCurrency, order)
	case int32(bitesla_srv_trader.TradeSide_BUY_MARKET):
		err = api.MarketBuy(ctx, reqCurrency, order)
	case int32(bitesla_srv_trader.TradeSide_SELL_MARKET):
		err = api.MarketSell(ctx, reqCurrency, order)
	default:
		return errors.New("订单类型不存在")
	}
	return err
}

func (e *exchangeResposity) CancelOrder(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, b *bitesla_srv_trader.Boolean) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	b.IsBool = false
	return api.CancelOrder(ctx, reqCurrency, b)
}

func (e *exchangeResposity) GetOneOrder(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	return api.GetOneOrder(ctx, reqCurrency, order)
}

func (e *exchangeResposity) GetUnfinishOrders(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, orders *bitesla_srv_trader.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	return api.GetUnfinishOrders(ctx, reqCurrency, orders)
}

func (e *exchangeResposity) GetOrderHistorys(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	return api.GetOrderHistorys(ctx, reqCurrency, order)
}

func (e *exchangeResposity) GetTicker(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, ticker *bitesla_srv_trader.Ticker) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	return api.GetTicker(ctx, reqCurrency, ticker)
}

func (e *exchangeResposity) GetDepth(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, depth *bitesla_srv_trader.Depth) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	return api.GetDepth(ctx, reqCurrency, depth)
}

func (e *exchangeResposity) GetTrades(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, trades *bitesla_srv_trader.Trades) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	return api.GetTrades(ctx, reqCurrency, trades)
}

func (e *exchangeResposity) GetExchangeName(ctx context.Context, reqCurrency *bitesla_srv_trader.ReqCurrency, name *bitesla_srv_trader.Str) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	return api.GetExchangeName(ctx, reqCurrency, name)
}
