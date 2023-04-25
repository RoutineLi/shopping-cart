package models

import (
	"gorm.io/gorm"
	"time"
)

type Device struct {
	gorm.Model
	Id             uint   `gorm:"column:id; type:int; autoincrement;" json:"identity"`
	UserId         uint   `gorm:"column:userid; type:int;" json:"userid"`
	Name           string `gorm:"column:name; type:varchar(50);" json:"name"`
	Key            string `gorm:"column:key; type:varchar(50);" json:"key"`
	Secret         string `gorm:"column:secret; type:varchar(50);" json:"secret"`
	LastOnlineTime int64  `gorm:"column:last_online_time; type:int(11);" json:"last_online_time"`
}

func (s Device) TableName() string {
	return "device_basic"
}

// GetDeviceList 获取设备列表
func GetDeviceList(name string) *gorm.DB {
	tx := DB.Model(new(Device)).Select("device_basic.id, device_basic.name," +
		"device_basic.key, device_basic.secret, ub.id userid, device_basic.last_online_time").
		Joins("LEFT JOIN user_basic ub ON ub.id = device_basic.userid")
	if name != "" {
		tx.Where("device_basic.name LIKE ?", "%"+name+"%")
	}

	return tx
}

// UpdateDeviceOnlineTime 更新设备上线时间
// [params] 设备key
// [return] error
func UpdateDeviceOnlineTime(deviceKey string) error {
	err := DB.Model(new(Device)).Where("`key` = ?", deviceKey).
		Update("last_online_time", time.Now().Unix()).Error
	return err
}
