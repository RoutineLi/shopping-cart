package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc   zrpc.RpcClientConf
	Kafka     kq.KqConf
	DeviceRpc zrpc.RpcClientConf
	//ImRpc     zrpc.RpcClientConf
}
