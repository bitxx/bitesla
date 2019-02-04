package model

type TokenAuth struct {
	ExName    string `json:"exName" example:"huobi.pro"`
	ApiKey    string `json:"apiKey" example:"自行输入"`
	ApiSecret string `json:"apiSecret" example:"自行输入"`
}
