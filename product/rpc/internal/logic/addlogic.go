package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
	"graduate_design/define"
	"graduate_design/models"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *product.AddRequest) (*product.AddResponse, error) {
	item := &models.Product{
		Name:          in.Name,
		Img:           in.Img,
		Price:         in.Price,
		Origin:        in.Origin,
		Brand:         in.Brand,
		Specification: in.Specification,
		ShelfLife:     in.ShelfLife,
		Description:   in.Description,
		Count:         int(in.Count),
		Type:          in.Type,
		Latitude:      in.Latitude,
		Longitude:     in.Longitude,
		Location:      in.Location,
	}
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err := l.svcCtx.DB.Create(item).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return err
		}
		return nil
	})

	threading.GoSafe(func() {
		//get new product id
		id := item.Id

		//Redis cache-in
		value, _ := json.Marshal(item)
		_, err = l.svcCtx.RedisClient.SetnxEx(strconv.Itoa(int(id))+"P", string(value), 30*60)
		if err != nil {
			return
		}
		//update stock cache
		//m := map[string]string{
		//	"total":  strconv.Itoa(data.Count),
		//	"secbuy": "0",
		//}
		//err = l.svcCtx.RedisClient.Hmset("stock:"+strconv.Itoa(int(id)), m)
		//if err != nil {
		//	return
		//}
		_, err = l.svcCtx.RedisClient.Sadd(define.ProIdsCache, strconv.Itoa(int(id)))
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		return &product.AddResponse{Status: false}, err
	}
	return &product.AddResponse{Status: true}, nil
}
