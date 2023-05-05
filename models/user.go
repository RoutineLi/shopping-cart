package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        uint   `gorm:"column:id; type:int; autoincrement;" json:"id"`
	Nickname  string `gorm:"column:nickname; type:varchar(50);" json:"username"`
	Password  string `gorm:"column:password; type:varchar(50);" json:"password"`
	Avatar    string `gorm:"column:avatar; type:varchar(512);" json:"avatar"`
	Motto     string `gorm:"column:motto; type:varchar(50);" json:"motto"`
	Gender    string `gorm:"column:gender; type:varchar(50);" json:"gender"`
	Age       uint   `gorm:"column:age; type:int;" json:"age"`
	Phone     string `gorm:"column:phone; type:varchar(30);" json:"phone"`
	Email     string `gorm:"column:email; type:varchar(50);" json:"email"`
	IsAdmin   uint   `gorm:"column:admin; type:int; default:0;" json:"is_admin"`
	AppKey    string `gorm:"column:app_key; type:varchar(50); default:null;" json:"app_key"`
	AppSecret string `gorm:"column:app_secret; type:varchar(50); default:null;" json:"app_secret"`
}

func (table User) TableName() string {
	return "user_basic"
}
