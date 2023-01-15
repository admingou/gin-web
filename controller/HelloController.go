package controller
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct{

}

//访问首页逻辑
func (hello *HelloController) Hello(cntext *gin.Context){
     cntext.IndentedJSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"message": "ok",
		"uri_path": "hello",
	 })
}