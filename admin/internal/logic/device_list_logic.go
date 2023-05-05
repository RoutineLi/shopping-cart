package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/device/types/device"
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
	list := make([]*types.Device, 0)
	resp = new(types.DeviceListResponse)
	count := 0
	rpcRsp, _ := l.svcCtx.RpcDevice.Ids(context.Background(), &device.IdsRequest{})
	if len(rpcRsp.Ids) == 0 {
		return resp, nil
	}
	for i := int((req.Page - 1) * req.Size); i < len(rpcRsp.Ids); i++ {
		if count == int(req.Size) {
			break
		}
		in := &device.DetailRequest{Id: rpcRsp.Ids[i]}
		out, _ := l.svcCtx.RpcDevice.Detail(context.Background(), in)
		item := &types.Device{
			Name:           out.Name,
			UserId:         out.Userid,
			Key:            out.Key,
			Secret:         out.Secret,
			LastOnlineTime: out.Lastonlinetime,
		}
		list = append(list, item)
		count++
	}
	resp.Count = int64(len(list))
	resp.List = list
	return resp, nil
}
