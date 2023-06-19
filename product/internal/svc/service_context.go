package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/product/internal/config"
	"graduate_design/product/rpc/productclient"
	"graduate_design/user/rpc/userclient"
)

type ServiceContext struct {
	Config   config.Config
	UserAuth userclient.User
	AuthResp *userclient.UserAuthResponse
	Product  productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserAuth: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		Product:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
