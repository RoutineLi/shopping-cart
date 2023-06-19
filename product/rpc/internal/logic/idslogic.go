package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/define"
	"graduate_design/models"
	"strconv"

	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdsLogic {
	return &IdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IdsLogic) Ids(in *product.IdsRequest) (*product.IdsResponse, error) {
	var pids []uint32
	items, _ := l.svcCtx.RedisClient.Smembers(define.ProIdsCache)
	if len(items) != 0 {
		for _, x := range items {
			item, _ := strconv.Atoi(x)
			pids = append(pids, uint32(item))
		}
		return &product.IdsResponse{Ids: pids}, nil
	}
	err := l.svcCtx.DB.Model(new(models.Product)).Select("id").Scan(&pids).Error
	if err != nil {
		return nil, err
	}
	threading.GoSafe(func() {
		for _, id := range pids {
			l.svcCtx.RedisClient.Sadd(define.ProIdsCache, strconv.Itoa(int(id)))
		}
	})
	return &product.IdsResponse{Ids: pids}, nil
}
