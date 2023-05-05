package logic

import (
	"context"
	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"
	"graduate_design/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryProductListLogic {
	return &CategoryProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryProductListLogic) CategoryProductList(req *types.CategroyProductListRequest) (resp *types.CategroyProductListResponse, err error) {
	resp = new(types.CategroyProductListResponse)
	rsp := new(product.IdsResponse)
	var pids []uint32
	rsp, err = l.svcCtx.Product.Ids(context.Background(), &product.IdsRequest{})
	if err != nil {
		resp.Status = "failure"
		resp.Code = 400
		resp.Message = err.Error()
		return resp, err
	}
	pids = rsp.Ids

	for _, id := range pids {
		in := &product.DetailRequest{Id: id}
		out, _ := l.svcCtx.Product.Detail(context.Background(), in)
		if out.Data.Type == req.Name {
			item := &types.ProductBasic{
				Id:    uint(out.Data.Id),
				Name:  out.Data.Name,
				Image: out.Data.Img,
			}
			resp.Data = append(resp.Data, item)
		}
	}
	resp.Message = "获取" + req.Name + "的商品成功"
	resp.Status = "success"
	resp.Code = 200
	return resp, nil
}
