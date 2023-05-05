package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DeviceRPC zrpc.RpcClientConf
	Mqtt      struct {
		Broker   string
		ClientID string
		Password string
	}
	RedisCli redis.RedisConf
}
