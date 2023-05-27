package initialize

import (
	"github.com/gin-gonic/gin"
	"shop_api/user_web/middlewares"
	"shop_api/user_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
