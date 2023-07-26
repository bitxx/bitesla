package handler

import (
	"context"
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
	return s.repo.getStrategyDetail(ctx, strategyInfoReq, strategyInfoResp)
}

func (s *StrategyHandler) ListStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfosResp *bitesla_srv_strategy.StrategyInfos) error {
	return s.repo.listStrategy(ctx, strategyInfoReq, strategyInfosResp)
}

func (s *StrategyHandler) PutStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	return s.repo.putStrategy(ctx, strategyInfoReq, strategyInfoResp)
}

func (s *StrategyHandler) DeleteStrategy(ctx context.Context, strategyInfoReq *bitesla_srv_strategy.StrategyInfo, strategyInfoResp *bitesla_srv_strategy.StrategyInfo) error {
	return s.repo.deleteStrategy(ctx, strategyInfoReq, strategyInfoResp)
}
