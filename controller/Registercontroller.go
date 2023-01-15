package controller
import (
	"net/http"
    "fmt"

	"github.com/gin-gonic/gin"

	"myoneapi/tools"
	"myoneapi/modes"
)

type RegisterUser struct{
	*modes.User
}

func (register *RegisterUser) Registers(context *gin.Context){
    // 1、获取参数
	 var us modes.User
	 if err := context.ShouldBindJSON(&us); err == nil{
		//2、验证传递参数
		//3、保存数据
		fmt.Println(us)
        tools.InsertData(us)
		context.IndentedJSON(http.StatusOK, gin.H{
			"code": 2000,
			"mesage": "添加成功",
			"request_uri": context.FullPath(),
		})
	 }else{
		context.IndentedJSON(http.StatusOK, gin.H{
			"code": 4001,
			"mesage": "参数缺失",
			"request_uri": context.FullPath(),
		})
	 }
	 
}