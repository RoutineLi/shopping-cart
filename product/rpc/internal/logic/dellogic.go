package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/models"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"
	"strconv"
)

type DelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogic) Del(in *product.DelRequest) (*product.DelResponse, error) {
	err := l.svcCtx.DB.Model(new(models.Product)).Where("id = ?", in.Id).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return &product.DelResponse{Status: false}, err
	}
	err = l.svcCtx.DB.Debug().Where("id = ?", in.Id).Delete(new(models.Product)).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return &product.DelResponse{Status: false}, err
	}

	//Redis cache-out
	threading.GoSafe(func() {
		_, err = l.svcCtx.RedisClient.Del(strconv.Itoa(int(in.Id)) + "P")
		if err != nil {
			logx.Error("[CACHE ERROR]: ", err)
			return
		}
	})

	return &product.DelResponse{Status: true}, nil
}
