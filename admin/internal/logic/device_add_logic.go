package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/device/types/device"
)

type DeviceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceAddLogic {
	return &DeviceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceAddLogic) DeviceAdd(req *types.DeviceAddRequest) (resp *types.DeviceAddResponse, err error) {
	resp = new(types.DeviceAddResponse)
	in := &device.AddRequest{Name: req.Name, Userid: uint32(req.UserId)}
	out, _ := l.svcCtx.RpcDevice.Add(context.Background(), in)
	if out.Status != true {
		resp.Code = 400
		resp.Msg = "failure"
		return resp, err
	}
	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
