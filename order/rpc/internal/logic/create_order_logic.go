package logic

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"graduate_design/models"
	"os"
	"sync/atomic"
	"time"

	"graduate_design/order/rpc/internal/svc"
	"graduate_design/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		oid := genOrderID(time.Now())
		o := &models.Orders{
			Id:     oid,
			UserId: uint(in.Uid),
		}
		tx.Create(o)
		oi := &models.OrderItem{
			OrderId: oid,
			UserId:  uint(in.Uid),
			ProId:   uint(in.Pid),
		}
		tx.Create(oi)

		return nil
	})

	if err != nil {
		logx.Error("[ORDER ERROR]: ", err)
		return &order.CreateOrderResponse{}, err
	}
	return &order.CreateOrderResponse{}, nil
}

var num int64

func genOrderID(t time.Time) string {
	s := t.Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
