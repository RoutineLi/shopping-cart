package logic

import (
	"context"

	"graduate_design/buy/rpc/internal/svc"
	"graduate_design/buy/rpc/types/buy"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetProductsLogic {
	return &SetProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetProductsLogic) SetProducts(in *buy.BuyProductsRequest) (*buy.BuyProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &buy.BuyProductsResponse{}, nil
}
