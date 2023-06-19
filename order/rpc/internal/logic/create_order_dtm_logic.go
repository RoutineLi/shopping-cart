package logic

import (
	"context"

	"graduate_design/order/rpc/internal/svc"
	"graduate_design/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderDTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMLogic {
	return &CreateOrderDTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderDTMLogic) CreateOrderDTM(in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &order.AddOrderResponse{}, nil
}
