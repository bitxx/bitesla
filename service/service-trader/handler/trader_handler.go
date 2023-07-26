package handler

import (
	"context"
)

type TraderHandler struct {
	repo *traderRepository
}

func NewTraderHandler() *TraderHandler {
	repository := &traderRepository{}
	handler := &TraderHandler{
		repo: repository,
	}
	return handler
}

func (t *TraderHandler) ListTrader(ctx context.Context, reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfos *bitesla_srv_trader.TraderInfos) error {
	return t.repo.listTrader(reqTraderInfo, respTraderInfos)
}

func (t *TraderHandler) PutTrader(ctx context.Context, reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	return t.repo.putTrader(reqTraderInfo, respTraderInfo)
}

func (t *TraderHandler) GetTraderDetail(ctx context.Context, reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	return t.repo.getTraderDetail(reqTraderInfo, respTraderInfo)
}

func (t *TraderHandler) DeleteTrader(ctx context.Context, reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	return t.repo.deleteTrader(reqTraderInfo, respTraderInfo)
}

func (t *TraderHandler) SwitchTrader(ctx context.Context, reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	return t.repo.switchTrader(reqTraderInfo, respTraderInfo)
}
func (t *TraderHandler) UpdateTraderStatus(ctx context.Context, reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	return t.repo.updateTraderStatus(reqTraderInfo, respTraderInfo)
}
