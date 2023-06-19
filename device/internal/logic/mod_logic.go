package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/models"
	"strconv"

	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"

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

func (l *ModLogic) Mod(in *device.ModRequest) (*device.ModResponse, error) {
	userid, _ := strconv.Atoi(in.Userid)
	err := l.svcCtx.DB.Debug().Where("id = ?", in.Id).Updates(&models.Device{
		UserId: uint(userid),
		Name:   in.Name,
	}).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return &device.ModResponse{Status: false}, err
	}
	threading.GoSafe(func() {
		l.svcCtx.RedisClient.Del(strconv.Itoa(int(in.Id)) + "D")
		l.svcCtx.RedisClient.Del(strconv.Itoa(userid) + "BY_USERID")
	})
	return &device.ModResponse{Status: true}, nil
}
