package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/product/rpc/types/product"
)

type ProductModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductModifyLogic {
	return &ProductModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductModifyLogic) ProductModify(req *types.ProductModifyRequest) (resp *types.ProductModifyResponse, err error) {
	resp = new(types.ProductModifyResponse)
	rpcReq := &product.ModRequest{Id: uint32(req.Id)}
	rpcReq.Data = &product.AddRequest{
		Name:          req.Data.Name,
		Img:           req.Data.Img,
		Price:         req.Data.Price,
		Origin:        req.Data.Origin,
		Brand:         req.Data.Brand,
		Specification: req.Data.Specification,
		ShelfLife:     req.Data.ShelfLife,
		Description:   req.Data.Description,
		Count:         int32(req.Data.Count),
		Type:          req.Data.Type,
		Latitude:      req.Data.Latitude,
		Longitude:     req.Data.Longitude,
		Location:      req.Data.Location,
	}

	rpcRsp, _ := l.svcCtx.RpcProduct.Mod(context.Background(), rpcReq)
	if rpcRsp.Status != true {
		resp.Code = 400
		resp.Msg = "failure"
		return resp, nil
	}

	resp.Code = 200
	resp.Msg = "success"
	return resp, nil
}
