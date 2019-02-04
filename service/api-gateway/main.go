package main

import (
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/common/util/cache"
	"github.com/jason-wj/bitesla/service/api-gateway/conf"
	"github.com/jason-wj/bitesla/service/api-gateway/db"
	_ "github.com/jason-wj/bitesla/service/api-gateway/docs"
	"github.com/jason-wj/bitesla/service/api-gateway/router"
	"github.com/micro/go-web"
)

const srcName = "bitesla.api.gateway"

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

// @title bitesla
// @version 2.0
// @description 暂无描述
// @termsOfService http://www.wjblog.com/

// @contact.name idea_wj@163.com
// @contact.url http://www.wjblog.com/
// @contact.email idea_wj

// @license.name Apache 2.0
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
		logger.Fatal("api-gateway初始化失败，失败信息：", err)
	}

	if err := webSrv.Run(); err != nil {
		logger.Fatal("api-gateway启动失败，失败信息：", err)
	}
}
