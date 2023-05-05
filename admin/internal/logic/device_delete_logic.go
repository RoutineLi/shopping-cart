package logic

import (
	"context"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/device/types/device"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteRequest) (resp *types.DeviceDeleteResponse, err error) {
	resp = new(types.DeviceDeleteResponse)
	id, _ := strconv.Atoi(req.Id)
	in := &device.DelRequest{Id: uint32(id)}
	out, _ := l.svcCtx.RpcDevice.Del(context.Background(), in)
	if out.Status != true {
		resp.Msg = "failure"
		resp.Code = 400
		return resp, nil
	}
	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
