package tools
import (
    "os"
	"bufio"
	"encoding/json"
)

//基础配置
type Config struct{
	AppName string  `json:"app_name"`
	AppMode string  `json:"app_mode"`
	AppHost string  `json:"app_host"`
	AppPort string  `json:"app_port"`
	*DataBase  `json:"database"`
}
//数据库配置
type DataBase struct{
   MsqlHost     string     `json:"mysql_host"`
   MsqlPort     string     `json:"mysql_port"`
   MsqlUser     string     `json:"mysql_user"`
   MsqlPassword string     `json:"mysql_password"`
   MsqlDb       string     `json:"mysql_db"`
   MsqlCharSet  string     `json:"mysql_charset"`
}

var _cfg *Config = nil
func  ParseConfig(path string) (*Config, error){
	//读取配置数据
	file,error := os.Open(path)
	if error != nil{
		panic(error)
	}
	//关闭读取配置
	defer file.Close()
	reader := bufio.NewReader(file)
    decoer := json.NewDecoder(reader)
	if err := decoer.Decode(&_cfg);err != nil{
		 return nil, err
	}
	return _cfg, nil
}