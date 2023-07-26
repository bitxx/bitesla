package db

import (
	"github.com/bitxx/bitesla/common/constants"
	"github.com/bitxx/bitesla/common/errs"
	"github.com/bitxx/bitesla/service/service-exchange/orm"
)

func AddOrUpdateExchange(currentUserId, exchangeId int64, name, descriptio string) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}

	strategy := &orm.ExchangeORM{
		UserId:      currentUserId,
		Id:          exchangeId,
		Name:        name,
		Description: descriptio,
		CreateUser:  currentUserId,
		UpdateUser:  currentUserId,
	}
	return db.Omit(constants.CreateTime, constants.UpdateTime).Save(strategy).Error
}

func GetExchangeList(size, page int32) (exchanges []orm.ExchangeORM, err error) {
	db := GetInstance().GetMysqlDB()
	if size <= 0 {
		size = 10
	}

	err = db.Model(orm.ExchangeORM{}).Select("id,description,name,create_time,update_time").Limit(size).Offset((page - 1) * size).Scan(&exchanges).Error
	return
}

func IsExchangeExist(exchangeId int64) bool {
	var count int
	db := GetInstance().GetMysqlDB()
	db.Model(&orm.ExchangeORM{Id: exchangeId}).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func GetExchangeDetail(exchangeID int64) (*orm.ExchangeORM, error) {
	db := GetInstance().GetMysqlDB()
	exchange := &orm.ExchangeORM{Id: exchangeID}
	err := db.Model(orm.ExchangeORM{}).First(exchange).Error
	return exchange, err
}
