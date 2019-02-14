package orm

import "time"

type StrategyORM struct {
	Id          int64     `orm:"id" json:"id"`                   // id
	UserId      int64     `orm:"user_id" json:"user_id"`         // 用户id
	StrategyId  int64     `orm:"strategy_id" json:"strategy_id"` // 策略id
	Description string    `orm:"description" json:"description"` // 策略描述
	Name        string    `orm:"name" json:"name"`               // 策略名，方便记忆
	Script      string    `orm:"script" json:"script"`           // 策略脚本
	CreateUser  int64     `orm:"create_user" json:"create_user"` // 创建人
	UpdateUser  int64     `orm:"update_user" json:"update_user"` // 更新人
	CreateTime  time.Time `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime  time.Time `orm:"update_time" json:"update_time"` // 更新时间
}

func (*StrategyORM) TableName() string {
	return "t_strategy"
}
