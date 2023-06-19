package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	Id         string  `gorm:"column:id;type:varchar(64);primaryKey" json:"id"`
	UserId     uint    `gorm:"column:userid;type:int;default:0" json:"userid"`
	ShoppingId uint    `gorm:"column:shopping_id;type:int;default:0" json:"shoppingid"`
	Payment    float64 `gorm:"column:payment;type:decimal(20,2);default:0" json:"payment"`
	PayType    int     `gorm:"column:pay_type;type:tinyint(4);default:1" json:"paytype"`
	Status     int     `gorm:"column:status;type:smallint(6);default:10" json:"status"`
	PayTime    string  `gorm:"column:pay_time;type:timestamp;default:null" json:"paytime"`
	EndTime    string  `gorm:"column:end_time;type:timestamp;default:null" json:"endtime"`
	CloseTime  string  `gorm:"column:close_time;type:timestamp;default:null" json:"closetime"`
}

func (s *Orders) TableName() string {
	return "orders_basic"
}

type OrderItem struct {
	gorm.Model
	Id       int64   `gorm:"column:id;type:bigint(20);PrimaryKey" json:"id"`
	OrderId  string  `gorm:"column:order_id;type:varchar(64);Key;" json:"orderid"`
	UserId   uint    `gorm:"column:user_id;type:int;default:0;Key" json:"userid"`
	ProName  string  `gorm:"column:pro_name;type:varchar(100);" json:"proname"`
	ProId    uint    `gorm:"column:pro_id;type:int;default:0;Key" json:"proid"`
	ProImage string  `gorm:"column:pro_image;type:varchar(500);" json:"proimage"`
	Price    float64 `gorm:"column:price;type:decimal(20,2);default:0" json:"price"`
	Count    int     `gorm:"column:count;type:int(10);default:0" json:"count"`
	Total    float64 `gorm:"column:total_price;type:decimal(20,2);default:0" json:"total"`
}

func (s *OrderItem) TableName() string {
	return "order_item_basic"
}

type Shopping struct {
	gorm.Model
	Id      int64  `gorm:"column:id;type:bigint(20);PrimaryKey" json:"id"`
	OrderId string `gorm:"column:order_id;type:varchar(64);Key;" json:"orderid"`
	UserId  uint   `gorm:"column:user_id;type:int;default:0;Key" json:"userid"`
}

func (s *Shopping) TableName() string {
	return "shopping_basic"
}
