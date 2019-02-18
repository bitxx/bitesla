package client

import (
	"context"
	"encoding/json"
	"github.com/jason-wj/bitesla/common/errs"
	pb "github.com/jason-wj/bitesla/service/service-trader/proto"
	"github.com/micro/go-micro/client"
)

type TraderClient struct {
	client pb.TraderService
}

func NewTraderClient() *TraderClient {
	c := pb.NewTraderService("", client.DefaultClient)
	return &TraderClient{
		client: c,
	}
}

//新增一个策略
func (client *TraderClient) PutTrader(data []byte) (interface{}, int, error) {
	traderInfo := &pb.TraderInfo{}
	err := json.Unmarshal(data, traderInfo)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if len(traderInfo.Name) <= 0 {
		return nil, errs.TraderNameErr, err
	}

	if len(traderInfo.Description) <= 0 {
		return nil, errs.TraderDescErr, err
	}

	if traderInfo.ExchangeId <= 0 {
		return nil, errs.TraderExchangeIdErr, err
	}

	if traderInfo.StrategyId <= 0 {
		return nil, errs.TraderStrategyIdErr, err
	}

	resp, err := client.client.PutTrader(context.Background(), traderInfo)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp, errs.Success, nil
}

//获取当前用户策略列表
func (client *TraderClient) ListTrader(data []byte) (interface{}, int, error) {
	traderInfo := &pb.TraderInfo{}
	err := json.Unmarshal(data, traderInfo)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if traderInfo.Page <= 0 {
		traderInfo.Page = 1
	}

	if traderInfo.Size <= 0 {
		traderInfo.Size = 10
	}

	resp, err := client.client.ListTrader(context.Background(), traderInfo)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp.TraderInfos, errs.Success, nil
}

func (client *TraderClient) GetTraderDetail(data []byte) (interface{}, int, error) {
	traderInfo := &pb.TraderInfo{}
	err := json.Unmarshal(data, traderInfo)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if traderInfo.TraderId <= 0 {
		return nil, errs.TraderIdErr, err
	}

	if traderInfo.ExchangeId <= 0 {
		return nil, errs.TraderExchangeIdErr, err
	}

	if traderInfo.StrategyId <= 0 {
		return nil, errs.TraderStrategyIdErr, err
	}

	resp, err := client.client.GetTraderDetail(context.Background(), traderInfo)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp, errs.Success, nil
}

func (client *TraderClient) SwitchTrader(data []byte, token string) (interface{}, int, error) {
	traderInfo := &pb.TraderInfo{}
	err := json.Unmarshal(data, traderInfo)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if traderInfo.TraderId <= 0 {
		return nil, errs.TraderIdErr, err
	}

	if traderInfo.ApiKey == "" || traderInfo.ApiSecret == "" {
		return nil, errs.ExchangeApiKeyAndSecret, nil
	}

	if len(token) <= 0 {
		return nil, errs.TokenErr, nil
	}

	traderInfo.Token = token

	resp, err := client.client.SwitchTrader(context.Background(), traderInfo)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp, errs.Success, nil
}
