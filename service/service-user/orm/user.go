package orm

import "time"

type UserORM struct {
	Id         int       `orm:"id" json:"id"`                   // id
	UserId     string    `orm:"user_id" json:"user_id"`         // 用户id
	Username   string    `orm:"username" json:"username"`       // 用户名
	Email      string    `orm:"email" json:"email"`             // 邮箱
	Password   string    `orm:"password" json:"password"`       // 密码
	Phone      string    `orm:"phone" json:"phone"`             // 手机号
	Nickname   string    `orm:"nickname" json:"nickname"`       // 昵称
	Birthday   string    `orm:"birthday" json:"birthday"`       // 生日
	Sex        int       `orm:"sex" json:"sex"`                 // 性别
	CreateUser string    `orm:"create_user" json:"create_user"` // 创建人
	UpdateUser string    `orm:"update_user" json:"update_user"` // 更新人
	CreateTime time.Time `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `orm:"update_time" json:"update_time"` // 更新时间
}

func (*UserORM) TableName() string {
	return "t_user"
}
