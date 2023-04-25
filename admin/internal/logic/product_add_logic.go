package logic

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/define"
	"graduate_design/models"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
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
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err = l.svcCtx.DB.Create(&models.Product{
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
		}).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			resp.Msg = err.Error()
			resp.Code = 400
			return err
		}
		//get new product id
		var id int64
		data := new(models.Product)
		err = l.svcCtx.DB.Model(new(models.Product)).Select("id").Where("name = ?", req.Data.Name).
			Scan(&id).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			resp.Msg = err.Error()
			resp.Code = 400
			return err
		}
		//Redis cache-in
		value, _ := json.Marshal(data)
		err = l.svcCtx.RedisClient.Hmset(define.ProCache, map[string]string{strconv.Itoa(int(id)): string(value)})
		if err != nil {
			logx.Error("[CACHE ERROR]: ", err)
			resp.Msg = err.Error()
			resp.Code = 400
			return err
		}
		return nil
	})
	if err != nil {
		return resp, err
	}
	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
