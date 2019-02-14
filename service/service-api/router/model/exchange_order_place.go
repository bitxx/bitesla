package model

//	OrderType:
//	0 表示limitBuy
//	1 表示limitSell
//	2 表示marketBuy
//	3 表示marketSell
//	AccountType:
//	1 表示point
//	2 表示splot
type OrderPlace struct {
	ExchangeId int64  `json:"exchangeId" example:"2358885120275906"`
	ApiKey     string `json:"apiKey" example:"自行输入"`
	ApiSecret  string `json:"apiSecret" example:"自行输入"`

	Amount       string `json:"amount" example:"0"`
	Price        string `json:"price" example:"0"`
	AccountType  int32  `json:"accountType" example:"2"`
	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
	OrderType    int32  `json:"orderType" example:"2"`
}
