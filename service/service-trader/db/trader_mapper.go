package db

import (
	"github.com/bitxx/bitesla/common/constants"
	"github.com/bitxx/bitesla/common/errs"
	"github.com/bitxx/bitesla/service/service-trader/orm"
)

func AddOrUpdateTrader(traderId, userId, strategyId, exchangeId int64, name, description string) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}

	trader := &orm.TraderORM{
		Id:          traderId,
		UserId:      userId,
		StrategyId:  strategyId,
		ExchangeId:  exchangeId,
		Name:        name,
		Description: description,
		CreateUser:  userId,
		UpdateUser:  userId,
	}
	return db.Omit(constants.CreateTime, constants.UpdateTime).Save(trader).Error
}

func GetTraderList(size, page int32, userId int64) (traders []orm.TraderORM, err error) {
	db := GetInstance().GetMysqlDB()
	if size <= 0 {
		size = 10
	}
	trader := &orm.TraderORM{
		UserId: userId,
	}
	err = db.Model(orm.TraderORM{}).Select("exchange_id,trader_id,strategy_id,description,name,status,create_time,update_time").Where(trader).Limit(size).Offset((page - 1) * size).Scan(&traders).Error
	return
}

func GetTraderDetail(userId, traderId int64) (*orm.TraderORM, error) {
	db := GetInstance().GetMysqlDB()
	trader := &orm.TraderORM{}
	err := db.Model(&orm.TraderORM{UserId: userId, Id: traderId}).First(trader).Error
	return trader, err
}

// 更新状态
func UpdateTraderStatus(userId, traderId int64, status int) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}
	err := db.Model(&orm.TraderORM{Id: traderId, UserId: userId}).Omit(constants.CreateTime, constants.UpdateTime).Update(map[string]interface{}{"status": status, "update_user": userId}).Error
	return err
}

func IsTraderExist(traderId, currentUserId int64) bool {
	var count int
	db := GetInstance().GetMysqlDB()
	db.Model(&orm.TraderORM{Id: traderId, UserId: currentUserId}).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
