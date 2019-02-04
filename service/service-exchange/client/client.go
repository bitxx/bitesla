package client

import (
	"context"
	"encoding/json"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/service/service-exchange/exchange"
	pb "github.com/jason-wj/bitesla/service/service-exchange/proto"
	"github.com/micro/go-micro/client"
)

type TraderClient struct {
	client pb.ExchangeService
}

func NewTraderClient() *TraderClient {
	c := pb.NewExchangeService("", client.DefaultClient)
	return &TraderClient{c}
}

func (client *TraderClient) OrderPlace(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	account := &pb.Order{}
	switch currencyReq.OrderType {
	case int32(pb.TradeSide_BUY):
		account, err = client.client.LimitBuy(context.Background(), currencyReq)
	case int32(pb.TradeSide_SELL):
		account, err = client.client.LimitSell(context.Background(), currencyReq)
	case int32(pb.TradeSide_BUY_MARKET):
		account, err = client.client.MarketBuy(context.Background(), currencyReq)
	case int32(pb.TradeSide_SELL_MARKET):
		account, err = client.client.MarketSell(context.Background(), currencyReq)
	default:
		return nil, errs.ExchangeAccountTypeErr, err
	}

	if err != nil {
		return nil, errs.ExchangeAccountTypeErr, err
	}
	return account, errs.Success, nil
}

func (client *TraderClient) CancelOrder(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if currencyReq.OrderId == "" {
		return nil, errs.ExchangeOrderIDErr, nil
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	boolean, err := client.client.CancelOrder(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return boolean, errs.ExchangeOrderIDErr, nil
}

func (client *TraderClient) GetOneOrder(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if currencyReq.OrderId == "" {
		return nil, errs.ExchangeOrderIDErr, nil
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	order, err := client.client.GetOneOrder(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}

	return order, errs.Success, nil
}

func (client *TraderClient) GetUnfinishOrders(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	orders, err := client.client.GetUnfinishOrders(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return orders, errs.Success, nil
}

func (client *TraderClient) GetOrderHistorys(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	if currencyReq.CurrentPage <= 0 {
		currencyReq.CurrentPage = 1
	}
	if currencyReq.PageSize <= 0 {
		currencyReq.PageSize = 10
	}
	orders, err := client.client.GetOrderHistorys(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return orders, errs.Success, nil
}

func (client *TraderClient) GetAccount(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}

	account, err := client.client.GetAccount(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return account, errs.Success, nil
}

func (client *TraderClient) GetTicker(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	ticker, err := client.client.GetTicker(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return ticker, errs.Success, nil
}

func (client *TraderClient) GetDepth(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	depth, err := client.client.GetDepth(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return depth, errs.Success, nil
}

func (client *TraderClient) GetKlineRecords(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if currencyReq.Period < 0 {
		return nil, errs.ExchangePeriodErr, nil
	}
	if currencyReq.Size <= 0 {
		return nil, errs.ExchangeSizeErr, nil
	}
	if currencyReq.Since <= 0 {
		return nil, errs.ExchangeSinceErr, nil
	}

	klines, err := client.client.GetKlineRecords(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return klines.Klines, errs.Success, nil
}

func (client *TraderClient) GetTrades(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	if currencyReq.Since <= 0 {
		return nil, errs.ExchangeSinceErr, nil
	}
	trades, err := client.client.GetTrades(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return trades.Trades, errs.Success, nil
}

func (client *TraderClient) GetExchangeName(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data)
	if code != errs.Success {
		return nil, code, err
	}
	name, err := client.client.GetExchangeName(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return name.Str, errs.Success, nil
}

func commonJudge(data []byte) (*pb.ReqCurrency, int, error) {
	currencyReq := &pb.ReqCurrency{}
	err := json.Unmarshal(data, currencyReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if currencyReq.ApiKey == "" || currencyReq.ApiSecret == "" {
		return nil, errs.ExchangeApiKeyAndSecret, nil
	}

	if currencyReq.ExName == "" {
		return nil, errs.ExchangeNameErr, nil
	}
	return currencyReq, errs.Success, nil
}
