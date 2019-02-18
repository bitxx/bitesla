package model

type OrderHistory struct {
	ExchangeId   int64  `json:"exchangeId" example:"2358885120275906"`
	ApiKey       string `json:"apiKey" example:"自行输入"`
	ApiSecret    string `json:"apiSecret" example:"自行输入"`
	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
	CurrentPage  int32  `json:"currentPage" example:"1"`
	PageSize     int32  `json:"pageSize" example:"1"`
}
