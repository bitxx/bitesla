package model

type TraderDetail struct {
	TraderId   int64 `json:"traderId" example:"948904443912"`
	StrategyId int64 `json:"strategyId" example:"23575003451411"`
	ExchangeId int64 `json:"exchangeId" example:"2358885120275906"`
}
