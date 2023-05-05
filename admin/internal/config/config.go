package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	DeviceRpc  zrpc.RpcClientConf
	Redis      redis.RedisConf
}
