package logic

import (
	"context"

	"graduate_design/websocket/im/internal/svc"
	"graduate_design/websocket/im/types/websocket"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecvMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecvMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecvMessageLogic {
	return &RecvMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecvMessageLogic) RecvMessage(in *websocket.RecvMessageRequest) (*websocket.RecvMessageResponse, error) {
	// todo: add your logic here and delete this line

	return &websocket.RecvMessageResponse{}, nil
}
