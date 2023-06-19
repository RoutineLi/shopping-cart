package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/models"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModLogic {
	return &ModLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModLogic) Mod(in *product.ModRequest) (*product.ModResponse, error) {
	//修改库存操作
	if in.Count != 0 {
		item := &models.Product{}
		err := l.svcCtx.DB.Debug().Where("id = ?", in.Id).First(item).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return &product.ModResponse{Status: false}, err
		}
		item.Count -= int(in.Count)
		err = l.svcCtx.DB.Updates(item).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return &product.ModResponse{Status: false}, err
		}
	} else {
		err := l.svcCtx.DB.Debug().Where("id = ?", in.Id).Updates(&models.Product{
			Name:          in.Data.Name,
			Img:           in.Data.Img,
			Price:         in.Data.Price,
			Origin:        in.Data.Origin,
			Brand:         in.Data.Brand,
			Specification: in.Data.Specification,
			ShelfLife:     in.Data.ShelfLife,
			Description:   in.Data.Description,
			Count:         int(in.Data.Count),
			Type:          in.Data.Type,
			Latitude:      in.Data.Latitude,
			Longitude:     in.Data.Longitude,
			Location:      in.Data.Location,
		}).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return &product.ModResponse{Status: false}, err
		}
	}
	//Redis cache-out
	threading.GoSafe(func() {
		_, err := l.svcCtx.RedisClient.Del(strconv.Itoa(int(in.Id)) + "P")
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return
		}
	})
	return &product.ModResponse{Status: true}, nil
}
