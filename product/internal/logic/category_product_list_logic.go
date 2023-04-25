package logic

import (
	"context"
	"graduate_design/models"
	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"

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
	list := make([]*types.ProductBasic, 0)
	resp = new(types.CategroyProductListResponse)
	//TODO
	err = models.GetProductListByCategory(req.Name).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Data = nil
		resp.Message = "获取" + req.Name + "的商品失败"
		resp.Code = 400
		resp.Status = "failure"
		return resp, err
	}
	resp.Data = list
	resp.Message = "获取" + req.Name + "的商品成功"
	resp.Status = "success"
	resp.Code = 200
	return resp, nil
}
