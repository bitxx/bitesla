package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason-wj/bitesla/service/api-gateway/conf"
	"github.com/jason-wj/bitesla/service/api-gateway/midware/jwt"
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
	if conf.CurrentConfig.Mode != conf.DevMode {
		router.Use(jwt.JWT())
	}

	userRouter(router)
	traderRouter(router)
	return router
}
