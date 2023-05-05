package logic

import (
	"context"
	"graduate_design/device/types/device"
	"strconv"

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
	rpcReq := &device.ModRequest{
		Id:     uint32(req.Id),
		Userid: strconv.Itoa(int(req.UserId)),
		Name:   req.Name,
	}
	rpcRsp, _ := l.svcCtx.RpcDevice.Mod(context.Background(), rpcReq)
	if rpcRsp.Status != true {
		resp.Code = 400
		resp.Msg = "failure"
		return resp, nil
	}

	resp.Code = 200
	resp.Msg = "success"
	return resp, nil
}
