package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"graduate_design/buy/internal/logic"
	"graduate_design/buy/internal/svc"
	"graduate_design/buy/internal/types"
)

func BuyServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BuyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBuyServiceLogic(r.Context(), svcCtx)
		resp, err := l.BuyService(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
