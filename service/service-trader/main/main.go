package main

import (
	"errors"
	"os"
	"strconv"
)

//此文件为用户编写策略需要依赖的
//执行时，Dockerfile会将其复制到容器中

func init() {
	Init()
}

func main() {
	if len(os.Args) != 5 {
		panic(errors.New("参数输入错误"))
	}
	currentUserIdStr := os.Args[1]
	exchangeIdStr := os.Args[2]
	apiKey := os.Args[3]
	apiSecret := os.Args[4]

	currentUserId, _ := strconv.ParseInt(currentUserIdStr, 10, 64)
	exchangeId, _ := strconv.ParseInt(exchangeIdStr, 10, 64)

	Run(currentUserId, exchangeId, apiKey, apiSecret)
}
