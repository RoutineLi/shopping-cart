package logic

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/define"
	"graduate_design/device/deviceclient"
	"graduate_design/websocket/internal/svc"
	"graduate_design/websocket/internal/types"
	"time"
)

type GetShoppingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShoppingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShoppingListLogic {
	return &GetShoppingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShoppingListLogic) GetShoppingList(req *types.GetShoppingListRequest, conn *websocket.Conn) (resp *types.GetShoppingListResponse, err error) {
	//1.通过Token里的userid去绑定指定购物车
	//2.监听指定channel获取mqtt数据
	//3.websocket发送对端
	userId := l.svcCtx.AuthResp.Extend["userid"]
	d, err := l.svcCtx.RpcDevice.DetailByUserId(l.ctx, &deviceclient.DetailByUserIdRequest{Userid: userId})
	if err != nil {
		conn.WriteMessage(1, []byte("info: this user are not bind device yet!"))
		return
	}
	_ = conn.WriteMessage(1, []byte("info: hello! "+l.svcCtx.AuthResp.Extend["name"]))

	_ = conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	cli := &define.Client{Conn: conn, Uid: userId}
	define.ClientMap.Store(cli, d.Key)

	for {
		_, _, err = conn.ReadMessage()
		if err != nil {
			// 检查是否是超时错误
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logx.Error("[Websocket Error]: client is unhealthy!")
			}
			err = conn.Close()
			if err != nil {
				return nil, err
			}
			break
		}
	}

	return
}
