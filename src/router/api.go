package router

import (
	"github.com/gin-gonic/gin"
	"goImPro-service/src/api/handler/auth"
	"goImPro-service/src/middleware"
)

var (
	login auth.AuthHandler
)

// RegisterApiRoutes 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	var api *gin.RouterGroup

	router.Use(middleware.Cors())

	api = router.Group("/api")
	{
		//登录
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", login.Login)           //登录
			authGroup.POST("/registered", login.Registered) //注册
		}

	}

}
