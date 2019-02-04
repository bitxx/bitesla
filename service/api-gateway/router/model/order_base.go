package model

type OrderBase struct {
	ExName    string `json:"exName" example:"huobi.pro"`
	ApiKey    string `json:"apiKey" example:"自行输入"`
	ApiSecret string `json:"apiSecret" example:"自行输入"`

	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
	orderId      string `json:"orderId" example:"2xxxxx"`
}
