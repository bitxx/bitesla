package handler

import (
	"context"
	"errors"
	"github.com/jason-wj/bitesla/service/service-exchange/proto"
)

type exchangeResposity struct {
}

func (e *exchangeResposity) getKlineRecords(reqCurrency *bitesla_srv_trader.ReqCurrency, kLines *bitesla_srv_trader.Klines) (err error) {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	err = api.GetKlineRecords(context.Background(), reqCurrency, kLines)
	return err
}

func (e *exchangeResposity) getAccount(reqCurrency *bitesla_srv_trader.ReqCurrency, account *bitesla_srv_trader.Accounts) (err error) {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	err = api.GetAccount(context.Background(), reqCurrency, account)
	return err
}

func (e *exchangeResposity) OrderPlace(reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}

	switch reqCurrency.OrderType {
	case int32(bitesla_srv_trader.TradeSide_BUY):
		err = api.LimitBuy(context.Background(), reqCurrency, order)
	case int32(bitesla_srv_trader.TradeSide_SELL):
		err = api.LimitSell(context.Background(), reqCurrency, order)
	case int32(bitesla_srv_trader.TradeSide_BUY_MARKET):
		err = api.MarketBuy(context.Background(), reqCurrency, order)
	case int32(bitesla_srv_trader.TradeSide_SELL_MARKET):
		err = api.MarketSell(context.Background(), reqCurrency, order)
	default:
		return errors.New("订单类型不存在")
	}
	return err
}

func (e *exchangeResposity) CancelOrder(reqCurrency *bitesla_srv_trader.ReqCurrency, b *bitesla_srv_trader.Boolean) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	b.IsBool = false
	err = api.CancelOrder(context.Background(), reqCurrency, b)
	return err
}

func (e *exchangeResposity) GetOneOrder(reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	err = api.GetOneOrder(context.Background(), reqCurrency, order)
	return err
}

func (e *exchangeResposity) GetUnfinishOrders(reqCurrency *bitesla_srv_trader.ReqCurrency, orders *bitesla_srv_trader.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	err = api.GetUnfinishOrders(context.Background(), reqCurrency, orders)
	return err
}

func (e *exchangeResposity) GetOrderHistorys(reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	err = api.GetOrderHistorys(context.Background(), reqCurrency, order)
	return err
}

func (e *exchangeResposity) GetTicker(reqCurrency *bitesla_srv_trader.ReqCurrency, ticker *bitesla_srv_trader.Ticker) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	err = api.GetTicker(context.Background(), reqCurrency, ticker)
	return err
}

func (e *exchangeResposity) GetDepth(reqCurrency *bitesla_srv_trader.ReqCurrency, depth *bitesla_srv_trader.Depth) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	err = api.GetDepth(context.Background(), reqCurrency, depth)
	return err
}

func (e *exchangeResposity) GetTrades(reqCurrency *bitesla_srv_trader.ReqCurrency, trades *bitesla_srv_trader.Trades) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	err = api.GetTrades(context.Background(), reqCurrency, trades)
	return err
}

func (e *exchangeResposity) GetExchangeName(reqCurrency *bitesla_srv_trader.ReqCurrency, name *bitesla_srv_trader.Str) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	err = api.GetExchangeName(context.Background(), reqCurrency, name)
	return err
}
