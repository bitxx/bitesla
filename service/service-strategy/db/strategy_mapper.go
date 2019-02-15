package db

import (
	"github.com/jason-wj/bitesla/common/constants"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/service/service-strategy/orm"
)

func AddOrUpdateStrategy(userId, strategyId int64, name, description, script string) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}

	strategy := &orm.StrategyORM{
		UserId:      userId,
		StrategyId:  strategyId,
		Name:        name,
		Description: description,
		Script:      script,
		CreateUser:  userId,
		UpdateUser:  userId,
	}
	return db.Omit(constants.CreateTime, constants.UpdateTime).Save(strategy).Error
}

func GetStrategyList(size, page int32, userID int64) (strategys []orm.StrategyORM, err error) {
	db := GetInstance().GetMysqlDB()
	if size <= 0 {
		size = 10
	}
	strategy := &orm.StrategyORM{
		UserId: userID,
	}

	err = db.Model(orm.StrategyORM{}).Select("strategy_id,description,name,create_time,update_time").Where(strategy).Limit(size).Offset((page - 1) * size).Scan(&strategys).Error
	return
}

func GetStrategyDetail(userID, strategyID int64) (*orm.StrategyORM, error) {
	db := GetInstance().GetMysqlDB()
	strategy := &orm.StrategyORM{}
	err := db.Model(orm.StrategyORM{}).Where(&orm.StrategyORM{UserId: userID, StrategyId: strategyID}).First(strategy).Error
	return strategy, err
}

func IsStrategyExist(strategyId, currentUserId int64) bool {
	var count int
	db := GetInstance().GetMysqlDB()
	db.Model(orm.StrategyORM{}).Where(&orm.StrategyORM{StrategyId: strategyId, UserId: currentUserId}).Count(&count)
	if count > 0 {
		return false
	}
	return true
}
