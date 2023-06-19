package svc

import (
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/device/deviceclient"
	"graduate_design/user/rpc/userclient"
	"graduate_design/websocket/im/im"
	"graduate_design/websocket/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	RpcUser   userclient.User
	RpcDevice deviceclient.Device
	RpcIm     im.IM
	AuthResp  *userclient.UserAuthResponse
	Kq        queue.MessageQueue
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcUser:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		RpcDevice: deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRpc)),
		//RpcIm: im.NewIM(zrpc.MustNewClient(c.ImRpc)),
	}
}
