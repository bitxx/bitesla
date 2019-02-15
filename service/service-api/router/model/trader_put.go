package model

type TraderPut struct {
	TraderId    int64  `json:"traderId" example:"948904443912"`
	StrategyId  int64  `json:"strategyId" example:"23575003451411"`
	ExchangeId  int64  `json:"exchangeId" example:"2358885120275906"`
	Name        string `json:"name" example:"我要执行这个策略"`
	Description string `json:"description" example:"选好交易所，选好策略，准备执行！"`
	Script      string `json:"script" example:"我是策略脚本，长长的。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。"`
}
