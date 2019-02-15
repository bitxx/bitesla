package orm

import "time"

type TraderORM struct {
	Id          int64     `orm:"id" json:"id"`                   // id
	TraderId    int64     `orm:"trader_id" json:"trader_id"`     // 策略执行的id
	UserId      int64     `orm:"user_id" json:"user_id"`         // 用户id
	StrategyId  int64     `orm:"strategy_id" json:"strategy_id"` // 策略id
	ExchangeId  int64     `orm:"exchange_id" json:"exchange_id"` // 交易所id
	Name        string    `orm:"name" json:"name"`               // 此交易名，用户自行设置，便于记忆
	Description string    `orm:"description" json:"description"` // 当前trader执行描述
	Status      int       `orm:"status" json:"status"`           // 0:停止，1:运行
	CreateUser  int64     `orm:"create_user" json:"create_user"` // 创建人
	UpdateUser  int64     `orm:"update_user" json:"update_user"` // 更新人
	CreateTime  time.Time `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime  time.Time `orm:"update_time" json:"update_time"` // 更新时间
}

func (*TraderORM) TableName() string {
	return "t_trader"
}
