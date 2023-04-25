package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/device/deviceclient"
	"graduate_design/device/types/device"
	"graduate_design/open/internal/config"
	"graduate_design/user/rpc/userclient"
)

type ServiceContext struct {
	Config    config.Config
	RpcDevice device.DeviceClient
	RpcUser   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcDevice: deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC)),
		RpcUser:   userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
