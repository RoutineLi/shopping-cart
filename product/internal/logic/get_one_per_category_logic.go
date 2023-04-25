package logic

import (
	"context"
	"gorm.io/gorm"
	"graduate_design/models"

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
	var datas []types.Product
	//TODO
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err = l.svcCtx.DB.Model(new(models.Product)).Select("DISTINCT type").Find(&datas).Error
		if err != nil {
			return err
		}
		for _, p := range datas {
			var product types.ProductBasic
			err = l.svcCtx.DB.Model(new(models.Product)).Where("type = ?", p.Type).Find(&product).
				Limit(1).Error
			if err != nil {
				return err
			}
			resp.Data = append(resp.Data, &product)
		}
		return nil
	})
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Code = 400
		resp.Status = "failure"
		resp.Message = err.Error()
		return resp, err
	}
	resp.Code = 200
	resp.Status = "success"
	resp.Message = "每个类别拿一个商品成功"
	return resp, nil
}
