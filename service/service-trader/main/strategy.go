package main

import "fmt"

//这是一个策略的测试

func Init() {
	fmt.Println("-----------init")
}

func Run(currentUserID, exchangeId int64, appKey, appSecret string) {
	fmt.Println("currentUserID:", currentUserID)
	fmt.Println("exchangeId:", exchangeId)
	fmt.Println("appKey:", appKey)
	fmt.Println("appSecret:", appSecret)
}
