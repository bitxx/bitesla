package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// -------------------- router start---------------------
func swaggerRouterNoAuth(router *gin.Engine) {
	router.GET(Swagger, wrapHandler())
}

// -------------------- router stop---------------------

const (
	Swagger = "/swagger/*any"
)

func wrapHandler() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
