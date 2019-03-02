package handler

import (
	"context"
	"errors"
	"github.com/jason-wj/bitesla/common/util/idgenerate"
	"github.com/jason-wj/bitesla/service/service-strategy/db"
	"github.com/jason-wj/bitesla/service/service-strategy/proto"
)

type strategyRepository struct {
}

func (s *strategyRepository) listStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfosResp *bitesla_srv_strategy.StrategyInfos) error {
	strategys, err := db.GetStrategyList(strategyInfoReq.Size, strategyInfoReq.Page, strategyInfoReq.CurrentLoginUserID)
	if err != nil {
		return err
	}
	for _, strategy := range strategys {
		tmp := &bitesla_srv_strategy.StrategyInfo{
			StrategyId:  strategy.Id,
			Description: strategy.Description,
			Name:        strategy.Name,
			Script:      strategy.Script,
			CreateTime:  strategy.CreateTime.Unix(),
			UpdateTime:  strategy.UpdateTime.Unix(),
		}
		strategyInfosResp.StrategyInfos = append(strategyInfosResp.StrategyInfos, tmp)
	}

	return nil
}

func (s *strategyRepository) getStrategyDetail(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	strategy, err := db.GetStrategyDetail(strategyInfoReq.CurrentLoginUserID, strategyInfoReq.StrategyId)
	strategyInfoResp.StrategyId = strategy.Id
	strategyInfoResp.Name = strategy.Name
	strategyInfoResp.Description = strategy.Description
	strategyInfoResp.Script = strategy.Script
	strategyInfoResp.CreateTime = strategy.CreateTime.Unix()
	strategyInfoResp.UpdateTime = strategy.UpdateTime.Unix()
	return err
}

func (s *strategyRepository) putStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	var err error
	strategyID := strategyInfoReq.StrategyId
	if strategyID <= 0 {
		strategyID, err = idgenerate.GetId()
		if err != nil {
			return errors.New("策略id生成失败")
		}
	} else {
		exist := db.IsStrategyExist(strategyID, strategyInfoReq.CurrentLoginUserID)
		if !exist {
			return errors.New("该策略不存在，请检查策略ID是否正确")
		}
	}

	err = db.AddOrUpdateStrategy(strategyInfoReq.Language, strategyInfoReq.CurrentLoginUserID, strategyID, strategyInfoReq.Name, strategyInfoReq.Description, strategyInfoReq.Script)
	return err
}

//TODO 暂时不考虑实现
func (s *strategyRepository) deleteStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	panic("implement me")
}
