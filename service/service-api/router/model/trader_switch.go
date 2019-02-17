package model

type TraderSwitch struct {
	TraderId   int64  `json:"traderId" example:"948904443912"`
	StrategyId int64  `json:"strategyId" example:"23575003451411"`
	ExchangeId int64  `json:"exchangeId" example:"2358885120275906"`
	ApiKey     string `json:"apiKey" example:"自行输入"`
	ApiSecret  string `json:"apiSecret" example:"自行输入"`
}
