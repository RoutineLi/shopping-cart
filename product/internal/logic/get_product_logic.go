package logic

import (
	"context"
	"encoding/json"
	"graduate_design/define"
	"graduate_design/models"
	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"

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
	claim := new(models.Product)
	resp = new(types.GetProductResponse)
	jsondata := ""
	jsondata, _ = l.svcCtx.RedisClient.Hget(define.ProCache, string(req.Id))
	if jsondata != "" {
		json.Unmarshal([]byte(jsondata), &resp.Data)
		resp.Status = "success"
		resp.Code = 200
		resp.Message = "获取商品详情成功"
		return resp, nil
	}
	err = l.svcCtx.DB.Model(new(models.Product)).Where("id = ?", req.Id).First(claim).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Status = "failure"
		resp.Message = "获取商品详情失败"
		resp.Code = 400
		return resp, err
	}
	resp.Data = types.Product{
		Id:            claim.Id,
		Name:          claim.Name,
		Img:           claim.Img,
		Price:         claim.Price,
		Origin:        claim.Origin,
		Brand:         claim.Brand,
		Specification: claim.Specification,
		ShelfLife:     claim.ShelfLife,
		Description:   claim.Description,
		Count:         claim.Count,
		Type:          claim.Type,
		Latitude:      claim.Latitude,
		Longitude:     claim.Longitude,
		Location:      claim.Location,
	}
	var temp []byte
	temp, _ = json.Marshal(resp.Data)
	l.svcCtx.RedisClient.Hmset(define.ProCache, map[string]string{string(req.Id): string(temp)})
	resp.Status = "success"
	resp.Code = 200
	resp.Message = "获取商品详情成功"
	return resp, nil
}
