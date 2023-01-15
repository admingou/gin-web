package main
import (
	"myoneapi/tools"
	"myoneapi/controller"
	"github.com/gin-gonic/gin"
	
)

func main(){
	cfg, error := tools.ParseConfig("./config/app.json")
	if error != nil{
		panic(error.Error())
	}
	app := gin.Default()
	registerRouter(app)
	//启动服务
	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}


//路由设置
func registerRouter(app *gin.Engine){
   new(controller.HelloController).Router(app)
}