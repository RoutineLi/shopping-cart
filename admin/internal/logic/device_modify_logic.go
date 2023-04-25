package logic

import (
	"context"
	"graduate_design/define"
	"graduate_design/models"

	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceModifyLogic {
	return &DeviceModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceModifyLogic) DeviceModify(req *types.DeviceModifyRequest) (resp *types.DeviceModifyResponse, err error) {
	resp = new(types.DeviceModifyResponse)
	device := new(models.Device)
	err = l.svcCtx.DB.Debug().Where("id = ?", req.Id).Updates(&models.Device{
		UserId: req.UserId,
		Name:   req.Name,
	}).Find(device).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Code = 400
		resp.Msg = err.Error()
		return resp, err
	}
	//Redis cache-out
	_, err = l.svcCtx.RedisClient.Hdel(define.DevCache, device.Key)
	if err != nil {
		logx.Error("[CACHE ERROR]: ", err)
		resp.Code = 400
		resp.Msg = err.Error()
		return resp, nil
	}
	resp.Code = 200
	resp.Msg = "success"
	return resp, nil
}
