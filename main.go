package main
import (
	"myoneapi/tools"
	"myoneapi/routes"
	"github.com/gin-gonic/gin"
	
)

func main(){
	//配置文件加载
	cfg, error := tools.ParseConfig("./config/app.json")
	if error != nil{
		panic(error.Error())
	}
	//gin框架初始化
	app := gin.Default()
    //路由注册
	routes.RegisterRouter(app)
	//启动服务
	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}