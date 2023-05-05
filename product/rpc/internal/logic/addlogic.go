package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
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
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err := l.svcCtx.DB.Create(&models.Product{
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
		}).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return err
		}
		return nil
	})

	threading.GoSafe(func() {
		//get new product id
		var id int64
		data := new(models.Product)
		err = l.svcCtx.DB.Model(new(models.Product)).Select("id").Where("name = ?", in.Name).
			Scan(&id).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return
		}
		//Redis cache-in
		value, _ := json.Marshal(data)
		err = l.svcCtx.RedisClient.Setex(strconv.Itoa(int(id))+"P", string(value), 30*60)
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
