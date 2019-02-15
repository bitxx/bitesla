package db

import (
	"github.com/jason-wj/bitesla/common/constants"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/service/service-trader/orm"
)

func AddOrUpdateTrader(traderId, userId, strategyId, exchangeId int64, name, description string) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}

	trader := &orm.TraderORM{
		UserId:      userId,
		StrategyId:  strategyId,
		TraderId:    traderId,
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

func GetTraderDetail(userId, exchangeId, strategyId, traderId int64) (*orm.TraderORM, error) {
	db := GetInstance().GetMysqlDB()
	trader := &orm.TraderORM{}
	err := db.Model(orm.TraderORM{}).Where(&orm.TraderORM{UserId: userId, TraderId: traderId, ExchangeId: exchangeId, StrategyId: strategyId}).First(trader).Error
	return trader, err
}

func IsTraderExist(traderId, exchangeId, strategyId, currentUserId int64) bool {
	var count int
	db := GetInstance().GetMysqlDB()
	db.Model(orm.TraderORM{}).Where(&orm.TraderORM{TraderId: traderId, ExchangeId: exchangeId, StrategyId: strategyId, UserId: currentUserId}).Count(&count)
	if count > 0 {
		return false
	}
	return true
}
