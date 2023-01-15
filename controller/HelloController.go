package controller
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct{

}

// func (hello *HelloController) Router(engine *gin.Engine){
// 	engine.GET("/hello", hello.Hello)
// }

//解析 /hello路由
func (hello *HelloController) Hello(cntext *gin.Context){
     cntext.IndentedJSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"message": "ok",
		"uri_path": "hello",
	 })
}