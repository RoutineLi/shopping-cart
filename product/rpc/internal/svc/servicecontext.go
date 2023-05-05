package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"graduate_design/models"
	"graduate_design/product/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB *gorm.DB
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config: c,
		DB: models.DB,
		RedisClient: redis.MustNewRedis(c.Redis),
	}
}
