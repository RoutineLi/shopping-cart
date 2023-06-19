package logic

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	types2 "graduate_design/websocket/im/types"
	"graduate_design/websocket/internal/svc"
	"graduate_design/websocket/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type Msg struct {
	Msg      string `json:"msg"`
	From     string `json:"from_id"`
	To       string `json:"to_id"`
	SendTime int64  `json:"send_time"`
}

type CustomerServiceChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomerServiceChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerServiceChatLogic {
	return &CustomerServiceChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CustomerServiceChatLogic) CustomerServiceChat1(conn *websocket.Conn, req *types.CustomerServiceChatRequest) (resp *types.CustomerServiceChatResponse, err error) {
	// TODO:Demo
	_ = conn.WriteJSON(l.svcCtx.AuthResp.Extend)
	for {
		_, data, _ := conn.ReadMessage()
		if data != nil {
			msg := &Msg{}
			err = json.Unmarshal(data, msg)
			if err != nil {
				conn.WriteMessage(1, []byte("invalid msg type"))
				return nil, err
			}
			sid, _ := strconv.Atoi(msg.To)
			if types2.SessionMap[uint(sid)].Client != nil {
				_ = types2.SessionMap[uint(sid)].Client.WriteJSON(msg)
			}
		}
	}
	//types2.SessionMap[userId].
	return
}
