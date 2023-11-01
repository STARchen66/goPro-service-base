package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goImPro-service/pkg/goroutine_pool"
	"goImPro-service/pkg/logger"
	"goImPro-service/pkg/mysql"
	"goImPro-service/pkg/redis"
	"goImPro-service/src/config"
	"goImPro-service/src/middleware"
	"goImPro-service/src/router"
)

//辅助程序

// LoadConfiguration 加载配置
func LoadConfiguration() {
	//初始化日志
	initLogger()

	redis.InitClient()

	mysql.InitDb()

	goroutine_pool.ConnectPool()

	// todo 其他逻辑

}

func Start() {
	engine := gin.Default()

	//启动业务协程

	engine.Use(middleware.Recover)

	setRoute(engine)
	//根路由测试
	engine.GET("/ping", pang)

	//SetMode根据输入字符串设置gin模式
	gin.SetMode(config.Conf.Server.Mode)

	//启动其他服务

	err := engine.Run(config.Conf.Server.Listen)
	if err != nil {
		fmt.Println(err)
	}
}

func pang(cxt *gin.Context) {
	println("pang")
}

// 初始化日志方法
func initLogger() {
	logger.InitLogger(
		config.Conf.Log.FileName,
		config.Conf.Log.MaxSize,
		config.Conf.Log.MaxBackup,
		config.Conf.Log.MaxAge,
		config.Conf.Log.Compress,
		config.Conf.Log.Type,
		config.Conf.Log.Level,
	)
}

// 注册路由方法
func setRoute(r *gin.Engine) {
	router.RegisterApiRoutes(r)
	//注册其它路由
	router.RegisterWsRouters(r)
}
