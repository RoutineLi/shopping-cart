package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/product/rpc/types/product"
)

type ProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddLogic {
	return &ProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAddLogic) ProductAdd(req *types.ProductAddRequest) (resp *types.ProductAddResponse, err error) {
	resp = new(types.ProductAddResponse)
	in := &product.AddRequest{
		Name:          req.Data.Name,
		Img:           req.Data.Img,
		Price:         req.Data.Price,
		Origin:        req.Data.Origin,
		Brand:         req.Data.Brand,
		Specification: req.Data.Specification,
		Description:   req.Data.Description,
		ShelfLife:     req.Data.ShelfLife,
		Count:         int32(req.Data.Count),
		Type:          req.Data.Type,
		Latitude:      req.Data.Latitude,
		Longitude:     req.Data.Longitude,
		Location:      req.Data.Location,
	}
	rpcRsp, _ := l.svcCtx.RpcProduct.Add(context.Background(), in)
	if rpcRsp.Status != true {
		resp.Msg = "failure"
		resp.Code = 400
		return resp, nil
	}
	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
