package logic

import (
	"context"
	"graduate_design/product/rpc/types/product"

	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequest) (resp *types.ProductListResponse, err error) {
	list := make([]*types.Product, 0)
	resp = new(types.ProductListResponse)
	count := 0
	rpcRsp, _ := l.svcCtx.RpcProduct.Ids(context.Background(), &product.IdsRequest{})
	if len(rpcRsp.Ids) == 0 {
		return resp, nil
	}

	for i := int((req.Page - 1) * req.Size); i < len(rpcRsp.Ids); i++ {
		if count == int(req.Size) {
			break
		}
		in := &product.DetailRequest{Id: rpcRsp.Ids[i]}
		out, _ := l.svcCtx.RpcProduct.Detail(context.Background(), in)
		item := &types.Product{
			Name:          out.Data.Name,
			Img:           out.Data.Img,
			Price:         out.Data.Price,
			Origin:        out.Data.Origin,
			Brand:         out.Data.Brand,
			Specification: out.Data.Specification,
			ShelfLife:     out.Data.ShelfLife,
			Description:   out.Data.Description,
			Count:         int(out.Data.Count),
			Type:          out.Data.Type,
			Latitude:      out.Data.Latitude,
			Longitude:     out.Data.Longitude,
			Location:      out.Data.Location,
		}
		list = append(list, item)
		count++
	}

	resp.Count = int64(len(list))
	resp.List = list
	return resp, nil
}
