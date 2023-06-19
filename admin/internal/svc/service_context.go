package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/admin/internal/config"
	"graduate_design/device/deviceclient"
	"graduate_design/product/rpc/productclient"
	"graduate_design/user/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	RpcUser    userclient.User
	AuthUser   *userclient.UserAuthResponse
	RpcProduct productclient.Product
	RpcDevice  deviceclient.Device
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		RpcUser:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		RpcProduct: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		RpcDevice:  deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRpc)),
	}
}
