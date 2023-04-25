package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"graduate_design/define"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := define.MysqlDsn + "graduate_design?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR]: ", err)
	}
	err = db.AutoMigrate(&Device{}, &Product{}, &User{})
	if err != nil {
		log.Fatalln("[DB ERROR]: ", err)
	}
	DB = db
}
