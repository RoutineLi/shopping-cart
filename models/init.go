package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"graduate_design/define"
	"log"
)

var DB *gorm.DB

func NewDB(DBName string) {
	dsn := define.MysqlDsn + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR]: ", err)
	}
	if DBName == define.DevDB {
		err = db.AutoMigrate(&Device{})
	} else if DBName == define.UserDB {
		err = db.AutoMigrate(&User{})
	} else if DBName == define.ProDB {
		err = db.AutoMigrate(&Product{})
	} else if DBName == define.OrderDB {
		err = db.AutoMigrate(&Orders{}, &OrderItem{}, &Shopping{})
	} else if DBName == define.PayDB {
		err = db.AutoMigrate(&Pay{})
	}

	if err != nil {
		log.Fatalln("[DB ERROR]: ", err)
	}
	DB = db
}
