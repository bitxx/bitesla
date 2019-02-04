package model

type OrderBase2 struct {
	ExName    string `json:"exName" example:"huobi.pro"`
	ApiKey    string `json:"apiKey" example:"自行输入"`
	ApiSecret string `json:"apiSecret" example:"自行输入"`

	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
}
