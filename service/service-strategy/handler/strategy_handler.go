package handler

import (
	"context"
	"github.com/jason-wj/bitesla/service/service-strategy/proto"
)

type StrategyHandler struct {
	repo *strategyRepository
}

func NewStrategyHandler() *StrategyHandler {
	repository := &strategyRepository{}
	handler := &StrategyHandler{
		repo: repository,
	}
	return handler
}

func (s *StrategyHandler) GetStrategyDetail(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	return s.repo.GetStrategyDetail(ctx, strategyInfoReq, strategyInfoResp)
}

func (s *StrategyHandler) ListStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfosResp *bitesla_srv_strategy.StrategyInfos) error {
	return s.repo.ListStrategy(ctx, strategyInfoReq, strategyInfosResp)
}

func (s *StrategyHandler) PutStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	return s.repo.PutStrategy(ctx, strategyInfoReq, strategyInfoResp)
}

func (s *StrategyHandler) DeleteStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	return s.repo.DeleteStrategy(ctx, strategyInfoReq, strategyInfoResp)
}
