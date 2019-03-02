package orm

import "time"

type ExchangeORM struct {
	Id          int64     `orm:"id" json:"id"`                   // 交易所id
	UserId      int64     `orm:"user_id" json:"user_id"`         // 用户id
	Description string    `orm:"description" json:"description"` // 交易所描述
	Name        string    `orm:"name" json:"name"`               // 交易所名称
	CreateUser  int64     `orm:"create_user" json:"create_user"` // 创建人
	UpdateUser  int64     `orm:"update_user" json:"update_user"` // 更新人
	CreateTime  time.Time `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime  time.Time `orm:"update_time" json:"update_time"` // 更新时间
}

func (*ExchangeORM) TableName() string {
	return "t_exchange"
}
