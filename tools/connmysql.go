package tools
import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"myoneapi/modes"
)

var db *gorm.DB = nil
var err error = nil

//返回数据库连接操作句柄
// func GetDB() *gorm.DB{
// 	return db
// }

func InitMysql(cfg *Config) *gorm.DB{
   args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
   cfg.MsqlUser,
   cfg.MsqlPassword,
   cfg.MsqlHost,
   cfg.MsqlPort,
   cfg.MsqlDb,
   cfg.MsqlCharSet,
   )
   db, err := gorm.Open("mysql", args)
   if err != nil{
	    panic("failed to connect database,err:" + err.Error())
   }
   //自动创建数据库
   db.AutoMigrate(new(modes.User))
   return db
} 

func InsertData(u modes.User){
	db.Create(&u)
}