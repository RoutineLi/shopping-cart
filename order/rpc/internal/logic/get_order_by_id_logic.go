package logic

import (
	"context"

	"graduate_design/order/rpc/internal/svc"
	"graduate_design/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIdLogic) GetOrderById(in *order.GetOrderByIdRequest) (*order.GetOrderByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &order.GetOrderByIdResponse{}, nil
}
