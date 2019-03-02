package db

import (
	"github.com/jason-wj/bitesla/common/constants"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/service/service-strategy/orm"
)

func AddOrUpdateStrategy(language int32, userId, strategyId int64, name, description, script string) error {
	db := GetInstance().GetMysqlDB()
	if db == nil {
		return errs.DBInitError
	}

	strategy := &orm.StrategyORM{
		UserId:      userId,
		Id:          strategyId,
		Name:        name,
		Description: description,
		Script:      script,
		CreateUser:  userId,
		UpdateUser:  userId,
		Language:    int(language),
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

	err = db.Model(orm.StrategyORM{}).Select("id,description,name,create_time,update_time").Where(strategy).Limit(size).Offset((page - 1) * size).Scan(&strategys).Error
	return
}

func GetStrategyDetail(userID, strategyID int64) (*orm.StrategyORM, error) {
	db := GetInstance().GetMysqlDB()
	strategy := &orm.StrategyORM{}
	err := db.Model(&orm.StrategyORM{UserId: userID, Id: strategyID}).First(strategy).Error
	return strategy, err
}

func IsStrategyExist(strategyId, currentUserId int64) bool {
	var count int
	db := GetInstance().GetMysqlDB()
	db.Model(&orm.StrategyORM{Id: strategyId, UserId: currentUserId}).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
