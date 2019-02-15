package model

type StrategyInfo struct {
	StrategyId  int64  `json:"strategyId" example:"0"`
	Name        string `json:"name" example:"我是一个策略名称"`
	Description string `json:"description" example:"我是策略描述"`
	Script      string `json:"script" example:"我是策略脚本，长长的。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。"`
}
