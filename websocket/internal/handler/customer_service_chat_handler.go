package handler

import (
	types2 "graduate_design/websocket/im/types"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"graduate_design/websocket/internal/logic"
	"graduate_design/websocket/internal/svc"
	"graduate_design/websocket/internal/types"
)

func CustomerServiceChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomerServiceChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCustomerServiceChatLogic(r.Context(), svcCtx)

		conn, err := types.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		userId, _ := strconv.Atoi(svcCtx.AuthResp.Extend["userid"])
		types2.SessionMap[uint(userId)] = &types2.SendData{
			Client: conn,
		}

		resp, err := l.CustomerServiceChat1(conn, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
