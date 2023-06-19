package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/buy/internal/config"
	"graduate_design/buy/rpc/buyclient"
	"graduate_design/user/rpc/userclient"
)

type ServiceContext struct {
	Config   config.Config
	AuthResp *userclient.UserAuthResponse
	UserAuth userclient.User
	RpcBuy   buyclient.Buy
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserAuth: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		RpcBuy:   buyclient.NewBuy(zrpc.MustNewClient(c.BuyRPC)),
	}
}
