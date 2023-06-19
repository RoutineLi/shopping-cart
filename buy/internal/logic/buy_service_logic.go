package logic

import (
	"context"
	"graduate_design/buy/rpc/buyclient"

	"graduate_design/buy/internal/svc"
	"graduate_design/buy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyServiceLogic {
	return &BuyServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyServiceLogic) BuyService(req *types.BuyRequest) (resp *types.BuyResponse, err error) {
	//TODO:test
	_, err = l.svcCtx.RpcBuy.SetOrder(l.ctx, &buyclient.SetOrderRequest{UserId: req.UserId, ProductId: req.ProductId})
	if err != nil {
		return nil, err
	}
	return &types.BuyResponse{
		Status:  "success",
		Code:    200,
		Message: "购买成功",
	}, nil
}
