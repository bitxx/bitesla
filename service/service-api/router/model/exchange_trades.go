package model

type Trades struct {
	ExchangeId int64 `json:"exchangeId" example:"2358885120275906"`

	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
	since        int32  `json:"since" example:"1"`
}
