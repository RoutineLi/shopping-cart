package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"graduate_design/websocket/internal/logic"
	"graduate_design/websocket/internal/svc"
	"graduate_design/websocket/internal/types"
)

func GetShoppingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetShoppingListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		conn, err := types.Upgrader.Upgrade(w, r, nil)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetShoppingListLogic(r.Context(), svcCtx)
		resp, err := l.GetShoppingList(&req, conn)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
