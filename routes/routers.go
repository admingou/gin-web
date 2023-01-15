package routes
import (
	 "myoneapi/controller"
	 "github.com/gin-gonic/gin"
)
type initRouter struct{
	controller.HelloController
}
//路由列表
func (initrouter *initRouter) Router(engine *gin.Engine){
	//页面组路由
	home := engine.Group("/home")
	//api组路由
	api := engine.Group("api")
	//api版本
	v1 := api.Group("v1")
	//首页
	home.GET("/", initrouter.Hello)
	//v1版本路由组
	v1.GET("/hello", initrouter.Hello)
}

//路由设置
func RegisterRouter(app *gin.Engine){ 
	new(initRouter).Router(app)
 }