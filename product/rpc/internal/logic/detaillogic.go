package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/models"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	metadata, _ := l.svcCtx.RedisClient.Get(strconv.Itoa(int(in.Id)) + "P")
	resp := new(product.DetailResponse)
	claim := new(models.Product)
	if metadata != "" {
		json.Unmarshal([]byte(metadata), claim)
		resp.Data = &product.Item{
			Id:            uint32(claim.Id),
			Name:          claim.Name,
			Img:           claim.Img,
			Price:         claim.Price,
			Origin:        claim.Origin,
			Brand:         claim.Brand,
			Specification: claim.Specification,
			ShelfLife:     claim.ShelfLife,
			Description:   claim.Description,
			Count:         int32(claim.Count),
			Type:          claim.Type,
			Latitude:      claim.Latitude,
			Longitude:     claim.Longitude,
			Location:      claim.Location,
		}
		return resp, nil
	}

	err := l.svcCtx.DB.Model(new(models.Product)).Where("id = ?", in.Id).First(claim).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return nil, err
	}
	resp.Data = &product.Item{
		Id:            uint32(claim.Id),
		Name:          claim.Name,
		Img:           claim.Img,
		Price:         claim.Price,
		Origin:        claim.Origin,
		Brand:         claim.Brand,
		Specification: claim.Specification,
		ShelfLife:     claim.ShelfLife,
		Description:   claim.Description,
		Count:         int32(claim.Count),
		Type:          claim.Type,
		Latitude:      claim.Latitude,
		Longitude:     claim.Longitude,
		Location:      claim.Location,
	}
	threading.GoSafe(func() {
		temp, _ := json.Marshal(resp.Data)
		l.svcCtx.RedisClient.Setex(strconv.Itoa(int(in.Id))+"P", string(temp), 30*60)
	})
	return resp, nil
}
