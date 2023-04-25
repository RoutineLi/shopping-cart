package logic

import (
	"context"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/define"
	"graduate_design/models"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
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
	var id int64
	err = l.svcCtx.DB.Debug().Where("name LIKE ?", "%"+req.Name+"%").Updates(&models.Product{
		Name:          req.Data.Name,
		Img:           req.Data.Img,
		Price:         req.Data.Price,
		Origin:        req.Data.Origin,
		Brand:         req.Data.Brand,
		Specification: req.Data.Specification,
		ShelfLife:     req.Data.ShelfLife,
		Description:   req.Data.Description,
		Count:         req.Data.Count,
		Type:          req.Data.Type,
		Latitude:      req.Data.Latitude,
		Longitude:     req.Data.Longitude,
		Location:      req.Data.Location,
	}).Scan(&id).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Code = 400
		resp.Msg = err.Error()
		return resp, err
	}
	//Redis cache-out
	_, err = l.svcCtx.RedisClient.Hdel(define.ProCache, strconv.Itoa(int(id)))
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Code = 400
		resp.Msg = err.Error()
		return resp, err
	}

	resp.Code = 200
	resp.Msg = "success"
	return resp, nil
}
