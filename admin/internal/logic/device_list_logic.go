package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
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
	resp = new(types.DeviceListResponse)
	count := 0
	rpcRsp, _ := l.svcCtx.RpcDevice.Ids(context.Background(), &device.IdsRequest{})
	if len(rpcRsp.Ids) == 0 {
		return resp, nil
	}
	var dids []uint32
	for i := int((req.Page - 1) * req.Size); i < len(rpcRsp.Ids); i++ {
		if count == int(req.Size) {
			break
		}
		dids = append(dids, rpcRsp.Ids[i])
		count++
	}
	ds, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, did := range dids {
			source <- did
		}
	}, func(item interface{}, writer mr.Writer[*types.Device], cancel func(error)) {
		id := item.(uint32)
		out, err := l.svcCtx.RpcDevice.Detail(l.ctx, &device.DetailRequest{Id: id})
		if err != nil {
			return
		}
		p := &types.Device{
			Name:           out.Name,
			UserId:         out.Userid,
			Key:            out.Key,
			Secret:         out.Secret,
			LastOnlineTime: out.Lastonlinetime,
		}
		writer.Write(p)
	}, func(pipe <-chan *types.Device, writer mr.Writer[[]*types.Device], cancel func(error)) {
		var r []*types.Device
		for p := range pipe {
			r = append(r, p)
		}
		writer.Write(r)
	})

	if err != nil {
		return nil, err
	}
	resp.List = ds
	resp.Count = int64(count)
	return resp, nil
}
