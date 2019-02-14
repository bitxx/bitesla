package model

type PhoneLogin struct {
	Password  string `json:"password" example:"sja123"`
	Phone     string `json:"phone" example:"13712345678"`
	PhoneCode string `json:"phoneCode" example:"9527"`
}
