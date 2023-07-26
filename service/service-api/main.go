package main

import (
	"github.com/bitxx/bitesla/common/logger"
	"github.com/bitxx/bitesla/common/util/cache"
	"github.com/bitxx/bitesla/service/service-api/conf"
	_ "github.com/bitxx/bitesla/service/service-api/docs"
	"github.com/bitxx/bitesla/service/service-api/router"
)

const srcName = "bitesla.service.api"

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
	issucc, err := mysql.GetInstance().InitPool()
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
}

// @title bitesla（比特斯拉）
// @version 2.0
// @description 代币量化交易框架，
// @termsOfService http://www.wjblog.top/

// @contact.name idea_wj@163.com
// @contact.url http://www.wjblog.top/
// @contact.email idea_wj

// @license.name Apache beta 1.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8090
// @BasePath /

// @securityDefinitions.apikey token
// @in header
// @name token
func main() {
	webSrv := web.NewService(
		web.Name(srcName),
		web.Address(":8090"),
	)

	r := router.GetAllRounters()

	webSrv.Handle("/", r)

	if err := webSrv.Init(); err != nil {
		logger.Fatal("service-api初始化失败，失败信息：", err)
	}

	if err := webSrv.Run(); err != nil {
		logger.Fatal("service-api启动失败，失败信息：", err)
	}
}
