package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/product/rpc/types/product"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteRequest) (resp *types.ProductDeleteResponse, err error) {
	resp = new(types.ProductDeleteResponse)
	in := &product.DelRequest{Id: uint32(req.Id)}
	rpcRsp, _ := l.svcCtx.RpcProduct.Del(context.Background(), in)
	if rpcRsp.Status != true {
		resp.Code = 400
		resp.Msg = "failure"
		return resp, nil
	}
	resp.Code = 200
	resp.Msg = "success"
	return resp, nil
}
