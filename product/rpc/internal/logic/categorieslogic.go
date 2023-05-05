package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
	"graduate_design/define"
	"graduate_design/models"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoriesLogic {
	return &CategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoriesLogic) Categories(in *product.CategoriesRequest) (*product.CategoriesResponse, error) {
	var categories []string
	categories, _ = l.svcCtx.RedisClient.Lrange(define.ProCache, 0, -1)
	if len(categories) != 0 {
		return &product.CategoriesResponse{Categories: categories}, nil
	}
	_ = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err := l.svcCtx.DB.Model(new(models.Product)).Select("DISTINCT type").Scan(&categories).Error
		if err != nil {
			return err
		}
		return nil
	})

	threading.GoSafe(func() {
		l.svcCtx.RedisClient.Lpush(define.ProCache, categories)
	})
	return &product.CategoriesResponse{Categories: categories}, nil
}
