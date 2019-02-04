package model

type OrderHistory struct {
	ExName       string `json:"exName" example:"huobi.pro"`
	ApiKey       string `json:"apiKey" example:"自行输入"`
	ApiSecret    string `json:"apiSecret" example:"自行输入"`
	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
	currentPage  int32  `json:"currentPage" example:"1"`
	pageSize     int32  `json:"pageSize" example:"1"`
}
