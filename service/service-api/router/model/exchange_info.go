package model

type ExchangeInfo struct {
	ExchangeId  int64  `json:"exchangeId" example:"0"`
	Name        string `json:"exName" example:"huobi.pro"`
	Description string `json:"description" example:"我是火币交易所"`
}
