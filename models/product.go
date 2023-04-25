package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id            uint    `gorm:"column:id; int; autoincrement;" json:"id,optional"`
	Name          string  `gorm:"column:name; varchar(50);" json:"name"`
	Img           string  `gorm:"column:img; varchar(50);" json:"img"`
	Price         float64 `gorm:"column:price; float;" json:"price"`
	Origin        string  `gorm:"column:origin; varchar(50);" json:"origin"`
	Brand         string  `gorm:"column:brand; varchar(20);" json:"brand"`
	Specification string  `gorm:"column:spec; varchar(30);" json:"specification"`
	ShelfLife     string  `gorm:"column:life; varchar(20);" json:"shelfLife"`
	Description   string  `gorm:"column:desc; varchar(50);" json:"description"`
	Count         int     `gorm:"column:count; int;" json:"count"`
	Type          string  `gorm:"column:type; varchar(20);" json:"type"`
	Latitude      float64 `gorm:"column:latitude; float;" json:"latitude"`
	Longitude     float64 `gorm:"column:longitude; float;" json:"longitude"`
	Location      string  `gorm:"column:location; varchar(50);" json:"location"`
}

func (s Product) TableName() string {
	return "product_basic"
}

// GetProductList 获取产品列表
func GetProductList(name string) *gorm.DB {
	tx := DB.Debug().Model(new(Product)).Select("id, name, img, price, origin, brand, spec, life, `desc`, count, type, latitude, longitude, location, created_at")
	if name != "" {
		tx.Where("name LIKE ?", "%"+name+"%")
	}
	return tx
}

// GetProductListByCategory 获取商品列表根据商品类别
func GetProductListByCategory(category string) *gorm.DB {
	tx := DB.Debug().Model(new(Product)).Select("id, name, img, price, origin, brand, spec, life, `desc`, count, type, latitude, longitude, location, created_at")
	if category != "" {
		tx.Where("type LIKE ?", "%"+category+"%")
	}
	return tx
}
