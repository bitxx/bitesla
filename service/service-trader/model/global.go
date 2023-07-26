package model

import (
	"github.com/bitxx/bitesla/service/service-strategy/client"
	"github.com/robertkrimen/otto"
)

type Global struct {
	Trader       *bitesla_srv_trader.TraderInfo
	StrategyInfo *bitesla_srv_trader.StrategyInfo
	Ctx          *otto.Otto
	Client       *client.StrategyClient
}

// js中的一个任务,目的是可以并发工作
type task struct {
	ctx  *otto.Otto    //js虚拟机
	fn   otto.Value    //代表该任务的js函数
	args []interface{} //函数的参数
}
