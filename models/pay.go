package models

import "gorm.io/gorm"

type Pay struct {
	gorm.Model
	Id         int64  `gorm:"column:id;type:bigint(20);PrimaryKey" json:"id"`
	OrderId    int64  `gorm:"column:order_id;type:varchar(64);Key;" json:"orderid"`
	UserId     uint   `gorm:"column:user_id;type:int;Key;default:0" json:"userid"`
	Platform   int    `gorm:"column:platform:type:tinyint(4)default:0" json:"platform"`
	PlatformId string `gorm:"column:platform_id;type:varchar(200);" json:"platformid"`
	Status     string `gorm:"column:status;type:varchar(20);" json:"status"`
}

func (s *Pay) TableName() string {
	return "pay_basic"
}
