package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/define"
	"graduate_design/models"
	"graduate_design/user/rpc/internal/svc"
	"graduate_design/user/rpc/types/user"
	"strconv"
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

func (l *IdsLogic) Ids(in *user.IdsRequest) (*user.IdsResponse, error) {
	var uids []uint32
	items, _ := l.svcCtx.RedisClient.Smembers(define.UserIdsCache)
	if len(items) != 0 {
		for _, x := range items {
			item, _ := strconv.Atoi(x)
			uids = append(uids, uint32(item))
		}
		return &user.IdsResponse{Ids: uids}, nil
	}
	err := l.svcCtx.DB.Model(new(models.User)).Select("id").Scan(&uids).Error
	if err != nil {
		return nil, err
	}

	threading.GoSafe(func() {
		for _, id := range uids {
			l.svcCtx.RedisClient.Sadd(define.UserIdsCache, strconv.Itoa(int(id)))
		}

	})

	return &user.IdsResponse{Ids: uids}, nil
}
