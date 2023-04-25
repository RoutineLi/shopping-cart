package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"graduate_design/models"
	"graduate_design/product/internal/config"
	"graduate_design/user/rpc/userclient"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	UserAuth    userclient.User
	AuthResp    *userclient.UserAuthResponse
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:      c,
		DB:          models.DB,
		UserAuth:    userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		RedisClient: redis.MustNewRedis(c.Redis),
	}
}
