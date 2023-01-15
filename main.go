package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"myoneapi/routes"
	"myoneapi/tools"
)

func main() {
	//配置文件加载
	cfg, error := tools.ParseConfig("./config/app.json")
	if error != nil {
		panic(error.Error())
	}
	fmt.Println(cfg.MsqlHost)
	//gin框架初始化
	app := gin.Default()
	//初始数据库
	db := tools.InitMysql(cfg)
	defer db.Close()       //延迟关闭数据库
	//路由注册
	routes.RegisterRouter(app)
	//启动服务
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
