package db

import (
	"github.com/jason-wj/bitesla/common/constants"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/service/service-user/orm"
)

//添加一个邮箱用户
func AddUserByEmail(email, password string) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}
	user := &orm.UserORM{
		Email:    email,
		Password: password,
	}
	db = db.Omit(constants.CreateTime, constants.UpdateTime, constants.CreateUser, constants.UpdateUser).Save(user)
	return db.Error
}

func LoginUserByEmail(email, pwd string) int {
	var count int
	db := GetInstance().GetMysqlDB()
	db.Model(orm.UserORM{}).Where(&orm.UserORM{Email: email, Password: pwd}).Count(&count)
	return count
}
