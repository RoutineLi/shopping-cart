package logic

import (
	"context"
	"graduate_design/websocket/im/internal/svc"
	"graduate_design/websocket/im/types"
	"graduate_design/websocket/im/types/websocket"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type Msg struct {
	Msg      string `json:"msg"`
	From     string `json:"from_id"`
	To       string `json:"to_id"`
	SendTime int64  `json:"send_time"`
}

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

func (l *SendMessageLogic) SendMessage(in *websocket.SendMessageRequest) (*websocket.SendMessageResponse, error) {
	ruid, _ := strconv.Atoi(in.Recvid)

	cli := types.SessionMap[uint(ruid)]
	data := &Msg{
		Msg:      in.Msg,
		From:     in.Sendid,
		To:       in.Recvid,
		SendTime: in.Sendtime,
	}
	//TODO
	if cli != nil {
		err := cli.Client.WriteJSON(data)
		if err != nil {
			return &websocket.SendMessageResponse{Status: "fail", Code: 400}, err
		}
	} else {
		//msg, _ := json.Marshal(data)
		//err := l.svcCtx.KafkaPusher.Push(string(msg))
		//if err != nil {
		return &websocket.SendMessageResponse{Status: "fail", Code: 400}, nil
		//}
	}

	return &websocket.SendMessageResponse{}, nil
}
