package db

import (
	"github.com/jason-wj/bitesla/common/constants"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/service/service-user/orm"
)

//添加一个邮箱用户
func AddUserByEmail(email, password string, id int64) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}
	user := &orm.UserORM{
		Id:       id,
		Email:    email,
		Password: password,
	}
	db = db.Omit(constants.CreateTime, constants.UpdateTime, constants.CreateUser, constants.UpdateUser).Save(user)
	return db.Error
}

func LoginUserByEmail(email, pwd string) (*orm.UserORM, error) {
	user := &orm.UserORM{}
	db := GetInstance().GetMysqlDB()
	err := db.Where(&orm.UserORM{Email: email, Password: pwd}).First(user).Error
	return user, err
}

func GetUserById(userId int64) (*orm.UserORM, error) {
	user := &orm.UserORM{}
	db := GetInstance().GetMysqlDB()
	err := db.Where(&orm.UserORM{Id: userId}).First(user).Error
	return user, err
}
