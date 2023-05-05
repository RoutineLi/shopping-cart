package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"graduate_design/admin/internal/config"
	"graduate_design/device/deviceclient"
	"graduate_design/models"
	"graduate_design/product/rpc/productclient"
	"graduate_design/user/rpc/userclient"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RpcUser     userclient.User
	AuthUser    *userclient.UserAuthResponse
	RpcProduct  productclient.Product
	RpcDevice   deviceclient.Device
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:      c,
		DB:          models.DB,
		RpcUser:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		RedisClient: redis.MustNewRedis(c.Redis),
		RpcProduct:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		RpcDevice:   deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRpc)),
	}
}
