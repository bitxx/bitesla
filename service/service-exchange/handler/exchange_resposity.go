package handler

import (
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
	kLines.Klines, err = api.GetKlineRecords(reqCurrency.CurrencyPair, reqCurrency.Period, reqCurrency.Size, reqCurrency.Since)
	return err
}

func (e *exchangeResposity) getAccount(reqCurrency *bitesla_srv_trader.ReqCurrency, account *bitesla_srv_trader.Accounts) (err error) {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	account.Accounts, err = api.GetAccount()
	return err
}

func (e *exchangeResposity) OrderPlace(reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}

	switch reqCurrency.OrderType {
	case int32(bitesla_srv_trader.TradeSide_BUY):
		order, err = api.LimitBuy(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, reqCurrency.AccountType)
	case int32(bitesla_srv_trader.TradeSide_SELL):
		order, err = api.LimitSell(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, reqCurrency.AccountType)
	case int32(bitesla_srv_trader.TradeSide_BUY_MARKET):
		order, err = api.MarketBuy(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, reqCurrency.AccountType)
	case int32(bitesla_srv_trader.TradeSide_SELL_MARKET):
		order, err = api.MarketSell(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, reqCurrency.AccountType)
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
	b.IsBool, err = api.CancelOrder(reqCurrency.OrderId, reqCurrency.CurrencyPair)
	return err
}

func (e *exchangeResposity) GetOneOrder(reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	order, err = api.GetOneOrder(reqCurrency.OrderId, reqCurrency.CurrencyPair)
	return err
}

func (e *exchangeResposity) GetUnfinishOrders(reqCurrency *bitesla_srv_trader.ReqCurrency, orders *bitesla_srv_trader.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	orders.Orders, err = api.GetUnfinishOrders(reqCurrency.CurrencyPair)
	return err
}

func (e *exchangeResposity) GetOrderHistorys(reqCurrency *bitesla_srv_trader.ReqCurrency, order *bitesla_srv_trader.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, true)
	if err != nil {
		return err
	}
	order.Orders, err = api.GetOrderHistorys(reqCurrency.CurrencyPair, reqCurrency.CurrentPage, reqCurrency.PageSize)
	return err
}

func (e *exchangeResposity) GetTicker(reqCurrency *bitesla_srv_trader.ReqCurrency, ticker *bitesla_srv_trader.Ticker) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	ticker, err = api.GetTicker(reqCurrency.CurrencyPair)
	return err
}

func (e *exchangeResposity) GetDepth(reqCurrency *bitesla_srv_trader.ReqCurrency, depth *bitesla_srv_trader.Depth) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	depth, err = api.GetDepth(reqCurrency.Size, reqCurrency.CurrencyPair)
	return err
}

func (e *exchangeResposity) GetTrades(reqCurrency *bitesla_srv_trader.ReqCurrency, trades *bitesla_srv_trader.Trades) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	trades.Trades, err = api.GetTrades(reqCurrency.CurrencyPair, reqCurrency.Since)
	return err
}

func (e *exchangeResposity) GetExchangeName(reqCurrency *bitesla_srv_trader.ReqCurrency, name *bitesla_srv_trader.Str) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExName, false)
	if err != nil {
		return err
	}
	name.Str = api.GetExchangeName()
	return nil
}
