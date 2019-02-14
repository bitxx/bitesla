package model

type Account struct {
	ExchangeId int64  `json:"exchangeId" example:"2358885120275906"`
	ApiKey     string `json:"apiKey" example:"自行输入"`
	ApiSecret  string `json:"apiSecret" example:"自行输入"`
}
