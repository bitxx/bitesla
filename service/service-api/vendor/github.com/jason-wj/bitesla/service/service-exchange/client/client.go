package client

import (
	"context"
	"encoding/json"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/service/service-exchange/exchange"
	pb "github.com/jason-wj/bitesla/service/service-exchange/proto"
	"github.com/micro/go-micro/client"
)

type ExchangeClient struct {
	client pb.ExchangeService
}

func NewExchangeClient() *ExchangeClient {
	c := pb.NewExchangeService("", client.DefaultClient)
	return &ExchangeClient{c}
}

func (client *ExchangeClient) OrderPlace(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, true)
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
		return nil, errs.Errors, err
	}
	return account, errs.Success, nil
}

func (client *ExchangeClient) CancelOrder(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, true)
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

func (client *ExchangeClient) GetOneOrder(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, true)
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

func (client *ExchangeClient) GetUnfinishOrders(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, true)
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

func (client *ExchangeClient) GetOrderHistorys(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, true)
	if code != errs.Success {
		return nil, code, err
	}
	if _, ok := exchange.CurrencyPair[currencyReq.CurrencyPair]; !ok {
		return nil, errs.ExchangeCoinErr, nil
	}
	if currencyReq.Page <= 0 {
		currencyReq.Page = 1
	}
	if currencyReq.Size <= 0 {
		currencyReq.Size = 10
	}
	orders, err := client.client.GetOrderHistorys(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return orders, errs.Success, nil
}

func (client *ExchangeClient) GetAccount(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, true)
	if code != errs.Success {
		return nil, code, err
	}

	account, err := client.client.GetAccount(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return account.Accounts, errs.Success, nil
}

func (client *ExchangeClient) GetTicker(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, false)
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

func (client *ExchangeClient) GetDepth(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, false)
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

func (client *ExchangeClient) GetKlineRecords(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, false)
	if code != errs.Success {
		return nil, code, err
	}
	if currencyReq.Period < 0 {
		return nil, errs.ExchangePeriodErr, nil
	}
	if currencyReq.Size <= 0 {
		currencyReq.Size = 10
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

func (client *ExchangeClient) GetTrades(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, false)
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

func (client *ExchangeClient) GetExchangeDetail(data []byte) (interface{}, int, error) {
	currencyReq, code, err := commonJudge(data, false)
	if code != errs.Success {
		return nil, code, err
	}

	exchange, err := client.client.GetExchangeDetail(context.Background(), currencyReq)
	if err != nil {
		return nil, errs.Errors, err
	}
	return exchange, errs.Success, nil
}

//新增一个策略
func (client *ExchangeClient) PutExchange(data []byte) (interface{}, int, error) {
	currency := &pb.Currency{}
	err := json.Unmarshal(data, currency)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if len(currency.ExName) <= 0 {
		return nil, errs.ExchangeNameErr, err
	}

	if len(currency.Description) <= 0 {
		return nil, errs.ExchangeDescriptionErr, err
	}

	resp, err := client.client.PutExchange(context.Background(), currency)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp, errs.Success, nil
}

//获取当前用户策略列表
func (client *ExchangeClient) ListExchange(data []byte) (interface{}, int, error) {
	currency := &pb.Currency{}
	err := json.Unmarshal(data, currency)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if currency.Page <= 0 {
		currency.Page = 1
	}

	if currency.Size <= 0 {
		currency.Size = 10
	}

	resp, err := client.client.ListExchange(context.Background(), currency)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp.Currencys, errs.Success, nil
}

//needAuth:是否需要key和secret
func commonJudge(data []byte, needAuth bool) (*pb.Currency, int, error) {
	currencyReq := &pb.Currency{}
	err := json.Unmarshal(data, currencyReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if needAuth && (currencyReq.ApiKey == "" || currencyReq.ApiSecret == "") {
		return nil, errs.ExchangeApiKeyAndSecret, nil
	}

	if currencyReq.ExchangeId <= 0 {
		return nil, errs.ExchangeIDErr, nil
	}
	return currencyReq, errs.Success, nil
}
