package logic

import (
	"context"
	"graduate_design/device/deviceclient"

	"graduate_design/open/internal/svc"
	"graduate_design/open/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageRequest) (resp *types.SendMessageResponse, err error) {
	l.svcCtx.RpcDevice.SendMessage(l.ctx, &deviceclient.SendMessageRequest{
		DeviceKey: req.DeviceKey,
		Msg:       req.Data,
	})
	if err != nil {
		logx.Error("[OPEN ERROR]: ", err.Error())
	}
	return
}
