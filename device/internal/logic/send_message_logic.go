package logic

import (
	"context"
	"errors"
	"fmt"
	"graduate_design/device/internal/mqtt"

	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageRequest) (*device.SendMessageResponse, error) {
	if in.DeviceKey == "" || in.Msg == "" {
		return &device.SendMessageResponse{
			Msg: errors.New("invalid params").Error(),
		}, errors.New("invalid params")
	}
	topic := "/sys/" + in.DeviceKey + "/receive"
	fmt.Println(topic)
	if token := mqtt.MC.Publish(topic, 0, false, in.Msg); token.Wait() && token.Error() != nil {
		logx.Error("[MQTT ERROR]: ", token.Error())
	}
	return &device.SendMessageResponse{
		Msg: "success",
	}, nil
}
