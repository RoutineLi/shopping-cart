package logic

import (
	"context"
	"graduate_design/models"
	"graduate_design/pkg"

	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceListLogic {
	return &DeviceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceListLogic) DeviceList(req *types.DeviceListRequest) (resp *types.DeviceListResponse, err error) {
	req.Size = pkg.If(req.Size == 0, 20, req.Size).(uint)
	req.Page = pkg.If(req.Page == 0, 0, (req.Page-1)*req.Size).(uint)
	var cnt int64
	resp = new(types.DeviceListResponse)
	data := make([]*types.Device, 0)
	err = models.GetDeviceList(req.Name).Count(&cnt).Limit(int(req.Size)).Offset(int(req.Page)).Find(&data).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return
	}
	resp.Count = cnt
	resp.List = data
	return resp, nil
}
