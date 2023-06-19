package svc

import (
	"gorm.io/gorm"
	"graduate_design/define"
	"graduate_design/models"
	"graduate_design/order/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB(define.OrderDB)
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
