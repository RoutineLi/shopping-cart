package logic

import (
	"context"
	"fmt"

	"graduate_design/order/rpc/internal/svc"
	"graduate_design/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderCheckLogic {
	return &CreateOrderCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderCheckLogic) CreateOrderCheck(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	fmt.Printf("CreateOrderCheck in:%v\n", in)
	return &order.CreateOrderResponse{}, nil
}
