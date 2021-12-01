package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoqingsong/studyweb/initial/config"
	"github.com/zhaoqingsong/studyweb/router"
	"github.com/zhaoqingsong/studyweb/service"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	service.Init()
	config.InitConfig()
	router.GetRouters(r)
	// GET：请求方式；/hello：请求的路径
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(config.Config.App.Address)
}