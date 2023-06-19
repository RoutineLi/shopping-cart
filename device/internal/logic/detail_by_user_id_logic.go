package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"
	"graduate_design/models"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailByUserIdLogic {
	return &DetailByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailByUserIdLogic) DetailByUserId(in *device.DetailByUserIdRequest) (*device.DetailByUserIdResponse, error) {
	item := &models.Device{}
	resp := &device.DetailByUserIdResponse{}
	metadata, _ := l.svcCtx.RedisClient.Get(in.Userid + "BY_USERID")
	if metadata != "" {
		json.Unmarshal([]byte(metadata), item)
		resp = &device.DetailByUserIdResponse{
			Id:             uint32(item.Id),
			Name:           item.Name,
			Userid:         strconv.Itoa(int(item.UserId)),
			Key:            item.Key,
			Secret:         item.Secret,
			Lastonlinetime: item.LastOnlineTime,
		}
		return resp, nil
	}
	id, _ := strconv.Atoi(in.Userid)
	err := l.svcCtx.DB.Model(new(models.Device)).Where("userid = ?", id).First(item).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return nil, err
	}
	resp = &device.DetailByUserIdResponse{
		Id:             uint32(item.Id),
		Name:           item.Name,
		Userid:         strconv.Itoa(int(item.UserId)),
		Key:            item.Key,
		Secret:         item.Secret,
		Lastonlinetime: item.LastOnlineTime,
	}

	threading.GoSafe(func() {
		temp, _ := json.Marshal(resp)
		l.svcCtx.RedisClient.SetnxEx(in.Userid+"BY_USERID", string(temp), 30*60)
	})
	return resp, nil
}
