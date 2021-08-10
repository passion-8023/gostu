package main

import (
	"fmt"
	"gostu/app/routers"
	"gostu/app/validates/rules"
	"gostu/pkg/config"
	"gostu/pkg/gin"
	"gostu/pkg/gorm"
	"gostu/pkg/logger"
	"log"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	//捕获 panic
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	//初始化gin框架
	srv := gin.Init()

	// 数据库初始化
	//if df := database.Init(); df != nil {
	//	defer df()
	//}
	if df := gorm.Init();df != nil {
		defer df()
	}

	//加载路由
	routers.Routers(srv)

	//初始化验证器
	rules.Init()

	//运行服务
	if err := srv.Run(config.AppConfig.GetString("http_port")); err != nil {
		fmt.Println("服务运行异常：", err)
		logger.Logger.Errorln("服务运行异常：", err)
	}
}
