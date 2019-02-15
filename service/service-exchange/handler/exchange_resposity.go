package handler

import (
	"errors"
	"github.com/jason-wj/bitesla/common/util/idgenerate"
	"github.com/jason-wj/bitesla/service/service-exchange/db"
	"github.com/jason-wj/bitesla/service/service-exchange/proto"
)

type exchangeResposity struct {
}

func (e *exchangeResposity) listExchange(currencyReq *bitesla_srv_exchange.Currency, currencyResps *bitesla_srv_exchange.Currencys) error {
	exchanges, err := db.GetExchangeList(currencyReq.Size, currencyReq.Page)
	if err != nil {
		return err
	}
	for _, exchange := range exchanges {
		tmp := &bitesla_srv_exchange.Currency{
			ExchangeId:  exchange.ExchangeId,
			Description: exchange.Description,
			ExName:      exchange.Name,
			CreateTime:  exchange.CreateTime.Unix(),
			UpdateTime:  exchange.UpdateTime.Unix(),
		}
		currencyResps.Currencys = append(currencyResps.Currencys, tmp)
	}
	return nil
}

func (e *exchangeResposity) putExchange(currencyReq *bitesla_srv_exchange.Currency, currencyResp *bitesla_srv_exchange.Currency) error {
	var err error
	exchangeId := currencyReq.ExchangeId
	if exchangeId <= 0 {
		exchangeId, err = idgenerate.GetId()
		if err != nil {
			return errors.New("交易所id生成失败")
		}
	} else {
		exist := db.IsExchangeExist(exchangeId)
		if !exist {
			return errors.New("该交易所不存在，请检查策略ID是否正确")
		}
	}

	err = db.AddOrUpdateExchange(currencyReq.CurrentLoginUserID, exchangeId, currencyReq.ExName, currencyReq.Description)
	return err
}

//TODO 暂时不考虑实现
func (e *exchangeResposity) deleteExchange(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Currency) error {
	panic("implement me")
}

func (e *exchangeResposity) getKlineRecords(reqCurrency *bitesla_srv_exchange.Currency, kLines *bitesla_srv_exchange.Klines) (err error) {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, false)
	if err != nil {
		return err
	}
	return api.GetKlineRecords(reqCurrency, kLines)
}

func (e *exchangeResposity) getAccount(reqCurrency *bitesla_srv_exchange.Currency, account *bitesla_srv_exchange.Accounts) (err error) {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, true)
	if err != nil {
		return err
	}
	return api.GetAccount(reqCurrency, account)
}

func (e *exchangeResposity) orderPlace(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, true)
	if err != nil {
		return err
	}

	switch reqCurrency.OrderType {
	case int32(bitesla_srv_exchange.TradeSide_BUY):
		err = api.LimitBuy(reqCurrency, order)
	case int32(bitesla_srv_exchange.TradeSide_SELL):
		err = api.LimitSell(reqCurrency, order)
	case int32(bitesla_srv_exchange.TradeSide_BUY_MARKET):
		err = api.MarketBuy(reqCurrency, order)
	case int32(bitesla_srv_exchange.TradeSide_SELL_MARKET):
		err = api.MarketSell(reqCurrency, order)
	default:
		return errors.New("订单类型不存在")
	}
	return err
}

func (e *exchangeResposity) cancelOrder(reqCurrency *bitesla_srv_exchange.Currency, b *bitesla_srv_exchange.Boolean) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, true)
	if err != nil {
		return err
	}
	b.IsBool = false
	return api.CancelOrder(reqCurrency, b)
}

func (e *exchangeResposity) getOneOrder(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, true)
	if err != nil {
		return err
	}
	return api.GetOneOrder(reqCurrency, order)
}

func (e *exchangeResposity) getUnfinishOrders(reqCurrency *bitesla_srv_exchange.Currency, orders *bitesla_srv_exchange.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, true)
	if err != nil {
		return err
	}
	return api.GetUnfinishOrders(reqCurrency, orders)
}

func (e *exchangeResposity) getOrderHistorys(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Orders) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, true)
	if err != nil {
		return err
	}
	return api.GetOrderHistorys(reqCurrency, order)
}

func (e *exchangeResposity) getTicker(reqCurrency *bitesla_srv_exchange.Currency, ticker *bitesla_srv_exchange.Ticker) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, false)
	if err != nil {
		return err
	}
	return api.GetTicker(reqCurrency, ticker)
}

func (e *exchangeResposity) getDepth(reqCurrency *bitesla_srv_exchange.Currency, depth *bitesla_srv_exchange.Depth) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, false)
	if err != nil {
		return err
	}
	return api.GetDepth(reqCurrency, depth)
}

func (e *exchangeResposity) getTrades(reqCurrency *bitesla_srv_exchange.Currency, trades *bitesla_srv_exchange.Trades) error {
	api, err := exchangeBuilder.APIKey(reqCurrency.ApiKey).APISecretkey(reqCurrency.ApiSecret).Build(reqCurrency.ExchangeId, false)
	if err != nil {
		return err
	}
	return api.GetTrades(reqCurrency, trades)
}

func (e *exchangeResposity) getExchangeDetail(reqCurrency *bitesla_srv_exchange.Currency, respCurrency *bitesla_srv_exchange.Currency) error {
	exchange, err := db.GetExchangeDetail(reqCurrency.ExchangeId)
	respCurrency.ExchangeId = exchange.ExchangeId
	respCurrency.ExName = exchange.Name
	respCurrency.CreateTime = exchange.CreateTime.Unix()
	respCurrency.UpdateTime = exchange.UpdateTime.Unix()
	respCurrency.Description = exchange.Description
	return err
}
