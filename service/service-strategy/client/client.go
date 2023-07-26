package client

import (
	"context"
	"encoding/json"
	"github.com/bitxx/bitesla/common/errs"
	pb "github.com/bitxx/bitesla/service/service-strategy/proto"
	"github.com/micro/go-micro/client"
)

type StrategyClient struct {
	client pb.StrategyService
}

func NewStrategyClient() *StrategyClient {
	c := pb.NewStrategyService("", client.DefaultClient)
	return &StrategyClient{
		client: c,
	}
}

// 新增一个策略
func (client *StrategyClient) PutStrategy(data []byte) (interface{}, int, error) {
	strategyInfo := &pb.StrategyInfo{}
	err := json.Unmarshal(data, strategyInfo)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if len(strategyInfo.Name) <= 0 {
		return nil, errs.StrategyNameErr, err
	}

	if len(strategyInfo.Description) <= 0 {
		return nil, errs.StrategyDescErr, err
	}

	if len(strategyInfo.Script) <= 0 {
		return nil, errs.StrategyScriptErr, err
	}

	_, ok := pb.Language_name[strategyInfo.Language]
	if !ok {
		return nil, errs.StrategyLanguageErr, err
	}

	resp, err := client.client.PutStrategy(context.Background(), strategyInfo)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp, errs.Success, nil
}

// 获取当前用户策略列表
func (client *StrategyClient) ListStrategy(data []byte) (interface{}, int, error) {
	strategyInfo := &pb.StrategyInfo{}
	err := json.Unmarshal(data, strategyInfo)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if strategyInfo.Page <= 0 {
		strategyInfo.Page = 1
	}

	if strategyInfo.Size <= 0 {
		strategyInfo.Size = 10
	}

	resp, err := client.client.ListStrategy(context.Background(), strategyInfo)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp.StrategyInfos, errs.Success, nil
}

func (client *StrategyClient) GetStrategyDetail(data []byte) (interface{}, int, error) {
	strategyInfo := &pb.StrategyInfo{}
	err := json.Unmarshal(data, strategyInfo)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}
	if strategyInfo.StrategyId <= 0 {
		return nil, errs.StrategyIdErr, err
	}

	resp, err := client.client.GetStrategyDetail(context.Background(), strategyInfo)
	if err != nil {
		return nil, errs.Errors, err
	}
	return resp, errs.Success, nil
}
