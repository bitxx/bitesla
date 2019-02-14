package model

type PhoneRegister struct {
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	PhoneCode string `json:"phoneCode"`
}
