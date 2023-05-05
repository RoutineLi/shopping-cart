package logic

import (
	"context"
	"graduate_design/product/rpc/types/product"

	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnePerCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnePerCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnePerCategoryLogic {
	return &GetOnePerCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnePerCategoryLogic) GetOnePerCategory(req *types.GetOnePerCategoryRequest) (resp *types.GetOnePerCategoryResponse, err error) {
	resp = new(types.GetOnePerCategoryResponse)
	cateRsp := new(product.CategoriesResponse)
	IdsRsp := new(product.IdsResponse)
	cateRsp, err = l.svcCtx.Product.Categories(context.Background(), &product.CategoriesRequest{})
	if err != nil {
		logx.Error("[RPC ERROR]: ", err)
		resp.Code = 400
		resp.Status = "failure"
		resp.Message = err.Error()
		return resp, err
	}
	IdsRsp, err = l.svcCtx.Product.Ids(context.Background(), &product.IdsRequest{})
	if err != nil {
		logx.Error("[RPC ERROR]: ", err)
		resp.Code = 400
		resp.Status = "failure"
		resp.Message = err.Error()
		return resp, err
	}
	for _, category := range cateRsp.Categories {
		for _, id := range IdsRsp.Ids {
			in := &product.DetailRequest{Id: id}
			out, _ := l.svcCtx.Product.Detail(context.Background(), in)
			if out.Data.Type == category {
				item := &types.ProductBasic{
					Id:    uint(out.Data.Id),
					Name:  out.Data.Name,
					Image: out.Data.Img,
				}
				resp.Data = append(resp.Data, item)
				break
			}
		}
	}
	resp.Code = 200
	resp.Status = "success"
	resp.Message = "每个类别拿一个商品成功"
	return resp, nil
}
