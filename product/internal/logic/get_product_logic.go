package logic

import (
	"context"
	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"
	"graduate_design/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductLogic) GetProduct(req *types.GetProductRequest) (resp *types.GetProductResponse, err error) {
	resp = new(types.GetProductResponse)
	rpcReq := &product.DetailRequest{Id: uint32(req.Id)}
	rpcRsp := &product.DetailResponse{}
	rpcRsp, err = l.svcCtx.Product.Detail(context.Background(), rpcReq)
	if err != nil {
		return nil, err
	}
	resp.Data = types.Product{
		Id:            uint(rpcRsp.Data.Id),
		Name:          rpcRsp.Data.Name,
		Img:           rpcRsp.Data.Img,
		Price:         rpcRsp.Data.Price,
		Origin:        rpcRsp.Data.Origin,
		Brand:         rpcRsp.Data.Brand,
		Specification: rpcRsp.Data.Specification,
		ShelfLife:     rpcRsp.Data.ShelfLife,
		Description:   rpcRsp.Data.Description,
		Count:         int(rpcRsp.Data.Count),
		Type:          rpcRsp.Data.Type,
		Latitude:      rpcRsp.Data.Latitude,
		Longitude:     rpcRsp.Data.Longitude,
		Location:      rpcRsp.Data.Location,
	}
	resp.Status = "success"
	resp.Code = 200
	resp.Message = "获取商品详情成功"
	return resp, nil
}
