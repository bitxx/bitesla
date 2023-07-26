package router

import (
	"github.com/bitxx/bitesla/service/service-api/conf"
	"github.com/bitxx/bitesla/service/service-api/midware/jwt"
	"github.com/gin-gonic/gin"
)

func GetAllRounters() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	if conf.CurrentConfig.Mode != conf.ProdMode {
		swaggerRouterNoAuth(router)
	}

	userRouterNoAuth(router)
	//非开发模式才用jwt
	/*if conf.CurrentConfig.Mode == conf.DevMode {
		router.Use(jwt.JWT())
	}*/
	router.Use(jwt.JWT())

	userRouter(router)
	exchangeRouter(router)
	strategyRouter(router)
	traderRouter(router)
	return router
}
