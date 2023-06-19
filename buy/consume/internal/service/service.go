package service

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/buy/consume/internal/config"
	"graduate_design/define"
	"graduate_design/order/rpc/orderclient"
	"graduate_design/product/rpc/productclient"
	"sync"
)

const (
	chanCount   = 10
	bufferCount = 1024
)

type Service struct {
	c          config.Config
	ProductRPC productclient.Product
	OrderRPC   orderclient.Order

	waiter   sync.WaitGroup
	msgsChan []chan *define.KBData
}

func NewService(c config.Config) *Service {
	s := &Service{
		c:          c,
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		msgsChan:   make([]chan *define.KBData, chanCount),
	}
	for i := 0; i < chanCount; i++ {
		ch := make(chan *define.KBData, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}
	return s
}

func (s *Service) consume(ch chan *define.KBData) {
	defer s.waiter.Done()

	for {
		m, ok := <-ch
		if !ok {
			logx.Error("consume service exit")
		}
		logx.Info("consume value: ", m)
		in := &productclient.ModRequest{
			Id:    uint32(m.Pid),
			Count: 1,
		}
		//_, err := s.ProductRPC.CheckAndUpdateStocks(context.Background(), &productclient.CAURequest{Pid: uint32(m.Pid)})
		//if err != nil {
		//	logx.Error("[ProductRPC ERROR]: ", err)
		//	return
		//}
		_, err := s.OrderRPC.CreateOrder(context.Background(), &orderclient.CreateOrderRequest{Uid: m.Uid, Pid: m.Pid})
		if err != nil {
			logx.Error("[OrderRPC ERROR]: ", err)
			return
		}
		_, err = s.ProductRPC.Mod(context.Background(), in)
		if err != nil {
			logx.Error("[ProductRPC ERROR]: ", err)
			return
		}

	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Info("consume value: ", value)
	var data []*define.KBData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	for _, d := range data {
		s.msgsChan[d.Pid%chanCount] <- d
	}
	return nil
}
