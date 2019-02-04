package model

// period:
//  0:  "1min",
//	1:  "5min",
//	2:  "15min",
//	3:  "30min",
//	4:  "60min",
//	5:  "1day",
//	6:  "1week",
//	7:  "1mon",
//	8:  "1year",

type Kline struct {
	ExName       string `json:"exName" example:"huobi.pro"`
	ApiKey       string `json:"apiKey" example:"自行输入"`
	ApiSecret    string `json:"apiSecret" example:"自行输入"`
	Size         int32  `json:"size" example:"150"`
	Period       int32  `json:"period" example:"0"`
	Since        int32  `json:"since" example:"10000000"`
	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
}
