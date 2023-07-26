package main

import (
	"github.com/bitxx/bitesla/common/logger"
	"github.com/bitxx/bitesla/common/util/cache"
	"github.com/bitxx/bitesla/common/util/idgenerate"
	"github.com/bitxx/bitesla/service/service-strategy/conf"
	"github.com/bitxx/bitesla/service/service-strategy/db"
	"github.com/bitxx/bitesla/service/service-strategy/handler"
	"time"
)

const (
	version = "lastest"
	srvName = "bitesla.srv.strategy"
)

func init() {
	conf.LoadConfig()
	isDebug := false

	//只有开发模式用窗口展示日志，其余模式都文本记录
	if conf.CurrentConfig.Mode == conf.DevMode {
		isDebug = true
	}

	//日志初始化
	logger.Init(isDebug, conf.CurrentConfig.LoggerConf.BaseFileName, conf.CurrentConfig.LoggerConf.LogLevel,
		conf.CurrentConfig.LoggerConf.EnableDynamic, conf.CurrentConfig.LoggerConf.JSONFormat,
		conf.CurrentConfig.LoggerConf.MaxAgeDays)

	//初始化mysql
	issucc, err := db.GetInstance().InitPool()
	if err != nil || !issucc {
		logger.Error(err)
		panic(err)
	}

	//初始化redis
	err = cache.Init(&cache.Cache{
		Url:         conf.CurrentConfig.Redis.Url,
		Password:    conf.CurrentConfig.Redis.Password,
		DBIndex:     conf.CurrentConfig.Redis.DbIndex,
		Key:         conf.CurrentConfig.Redis.DefaultKey,
		MaxIdle:     conf.CurrentConfig.Redis.MaxIdle,
		MaxActive:   conf.CurrentConfig.Redis.MaxActive,
		IdleTimeout: conf.CurrentConfig.Redis.IdleTimeout,
	})
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	redisCache, err := cache.GetRedisCache()
	err = redisCache.ClearAll()

	//id生成器
	idgenerate.Init(true)
}

func main() {
	strategyHandler := handler.NewStrategyHandler()

	// New Service
	service := micro.NewService(
		micro.Name(srvName),
		micro.Version(version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := bitesla_srv_strategy.RegisterStrategyHandler(service.Server(), strategyHandler)
	if err != nil {
		logger.Error("RegisterStrategyHandler err :", err)
		return
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
