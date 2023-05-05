package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"graduate_design/models"
	"graduate_design/user/internal/config"
	"graduate_design/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RpcUser userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
