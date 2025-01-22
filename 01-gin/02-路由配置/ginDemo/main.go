package main

import (
	"fmt"
	"ginDemo/config"
	"ginDemo/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 发布模式会关闭调试信息和日志输出，从而提升性能
	//gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式

	engine := gin.Default()

	router.InitRouter(engine) // 设置路由

	err := engine.Run(config.PORT)

	if err != nil {
		fmt.Println("server start failed")
		log.Fatal(err)
		return
	}

	fmt.Println("服务启动成功 端口: ", config.PORT)
}
