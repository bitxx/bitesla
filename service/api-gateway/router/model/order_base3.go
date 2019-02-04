package model

type OrderBase3 struct {
	ExName    string `json:"exName" example:"huobi.pro"`
	ApiKey    string `json:"apiKey" example:"自行输入"`
	ApiSecret string `json:"apiSecret" example:"自行输入"`

	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
	Size         int32  `json:"size" example:"1"`
}
