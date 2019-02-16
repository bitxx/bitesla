package main

import (
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/common/util/cache"
	"github.com/jason-wj/bitesla/common/util/idgenerate"
	"github.com/jason-wj/bitesla/service/service-trader/conf"
	"github.com/jason-wj/bitesla/service/service-trader/db"
	"github.com/jason-wj/bitesla/service/service-trader/handler"
	"github.com/jason-wj/bitesla/service/service-trader/proto"
	"github.com/jason-wj/bitesla/service/service-trader/trader"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"time"
)

const (
	version = "lastest"
	srvName = "bitesla.srv.trader"
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
	//初始化用户交易消费者
	err := trader.InitTraderQueue(conf.CurrentConfig.Nsq.TopicDefaultName, conf.CurrentConfig.Nsq.ChannelDefaultName, conf.CurrentConfig.Nsq.Tcp)
	if err != nil {
		logger.Error(err)
		return
	}

	traderHandler := handler.NewTraderHandler()

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
	err = bitesla_srv_trader.RegisterTraderHandler(service.Server(), traderHandler)
	if err != nil {
		logger.Error("RegisterTraderHandler err :", err)
		return
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
