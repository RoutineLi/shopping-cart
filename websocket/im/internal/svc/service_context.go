package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"graduate_design/websocket/im/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.Topic),
	}
}
