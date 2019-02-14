package model

type Ticker struct {
	ExchangeId   int64  `json:"exchangeId" example:"2358885120275906"`
	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
}
