package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"graduate_design/define"
	"graduate_design/device/internal/config"
	"graduate_design/models"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RedisClient *redis.Redis
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB(define.DevDB)
	return &ServiceContext{
		Config:      c,
		DB:          models.DB,
		RedisClient: redis.MustNewRedis(c.RedisCli),
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.Topic),
	}
}
