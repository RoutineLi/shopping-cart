package logic

import (
	"context"
	"graduate_design/models"
	"graduate_design/pkg"

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
	req.Size = pkg.If(req.Size == 0, 20, req.Size).(uint)
	req.Page = pkg.If(req.Page == 0, 0, (req.Page-1)*req.Size).(uint)
	resp = new(types.ProductListResponse)
	var count int64
	err = models.GetProductList(req.Name).Count(&count).Offset(int(req.Page)).Limit(int(req.Size)).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return
	}
	resp.Count = count
	resp.List = list
	return resp, nil
}
