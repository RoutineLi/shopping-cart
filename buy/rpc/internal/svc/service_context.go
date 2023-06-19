package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/buy/rpc/internal/config"
	"graduate_design/product/rpc/productclient"
)

type ServiceContext struct {
	Config      config.Config
	ProductRPC  productclient.Product
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ProductRPC:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.Topic),
	}
}
