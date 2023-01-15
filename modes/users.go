package modes
import (
	_ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)
type User struct{
	gorm.Model
    Phone       string        `gorm:"column:phone;varchar(11)" form:"phone" json:"phone" validate:"eq=11"`
	Name        string        `gorm:"column:name;varchar(20)" fo" form:"name" json:"name" validate:"gt=0;lt=20"`
	Password    string        `gorm:"column:password;varchar(20)" fo" form:"password" json:"password" validate:"gt=5;lt=15"`
	Address     string        `gorm:"column:address;varchar(255)" form:"address" json:"address" validate:"omitempty"`
	Role        int           `gorm:"column:role;int;default:2" form:"role" json:"role" validate:"omitempty"`
	Activation  int           `gorm:"column:activation;int;default:1" form:"activation" json:"activation" validate:"omitempty"`
}

