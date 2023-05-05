package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/define"
	"graduate_design/models"
	"strconv"

	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"

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

func (l *IdsLogic) Ids(in *device.IdsRequest) (*device.IdsResponse, error) {
	resp := &device.IdsResponse{}
	var dids []uint32
	items, _ := l.svcCtx.RedisClient.Lrange(define.DevIdsCache, 0, -1)
	if len(items) != 0 {
		for _, x := range items {
			item, _ := strconv.Atoi(x)
			dids = append(dids, uint32(item))
		}
		resp.Ids = dids
		return resp, nil
	}
	err := l.svcCtx.DB.Model(new(models.Device)).Select("id").Scan(&dids).Error
	if err != nil {
		return nil, err
	}
	threading.GoSafe(func() {
		l.svcCtx.RedisClient.Lpush(define.DevIdsCache, dids)
	})
	resp.Ids = dids
	return resp, nil
}
