package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/models"
	"strconv"

	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"

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

func (l *DetailLogic) Detail(in *device.DetailRequest) (*device.DetailResponse, error) {
	item := &models.Device{}
	resp := &device.DetailResponse{}
	metadata, _ := l.svcCtx.RedisClient.Get(strconv.Itoa(int(in.Id)) + "D")
	if metadata != "" {
		json.Unmarshal([]byte(metadata), item)
		resp = &device.DetailResponse{
			Id:             uint32(item.Id),
			Name:           item.Name,
			Userid:         strconv.Itoa(int(item.UserId)),
			Key:            item.Key,
			Secret:         item.Secret,
			Lastonlinetime: item.LastOnlineTime,
		}
		return resp, nil
	}
	err := l.svcCtx.DB.Model(new(models.Device)).Where("id = ?", in.Id).First(item).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return nil, err
	}
	resp = &device.DetailResponse{
		Id:             uint32(item.Id),
		Name:           item.Name,
		Userid:         strconv.Itoa(int(item.UserId)),
		Key:            item.Key,
		Secret:         item.Secret,
		Lastonlinetime: item.LastOnlineTime,
	}

	threading.GoSafe(func() {
		temp, _ := json.Marshal(resp)
		l.svcCtx.RedisClient.SetnxEx(strconv.Itoa(int(in.Id))+"D", string(temp), 30*60)
	})
	return resp, nil
}
