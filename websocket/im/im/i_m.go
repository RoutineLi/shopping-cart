// Code generated by goctl. DO NOT EDIT.
// Source: im.proto

package im

import (
	"context"

	"graduate_design/websocket/im/types/websocket"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	RecvMessageRequest  = websocket.RecvMessageRequest
	RecvMessageResponse = websocket.RecvMessageResponse
	SendMessageRequest  = websocket.SendMessageRequest
	SendMessageResponse = websocket.SendMessageResponse

	IM interface {
		SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
		RecvMessage(ctx context.Context, in *RecvMessageRequest, opts ...grpc.CallOption) (*RecvMessageResponse, error)
	}

	defaultIM struct {
		cli zrpc.Client
	}
)

func NewIM(cli zrpc.Client) IM {
	return &defaultIM{
		cli: cli,
	}
}

func (m *defaultIM) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	client := websocket.NewIMClient(m.cli.Conn())
	return client.SendMessage(ctx, in, opts...)
}

func (m *defaultIM) RecvMessage(ctx context.Context, in *RecvMessageRequest, opts ...grpc.CallOption) (*RecvMessageResponse, error) {
	client := websocket.NewIMClient(m.cli.Conn())
	return client.RecvMessage(ctx, in, opts...)
}