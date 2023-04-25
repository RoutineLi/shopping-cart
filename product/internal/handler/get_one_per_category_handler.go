package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"graduate_design/product/internal/logic"
	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"
)

func GetOnePerCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOnePerCategoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetOnePerCategoryLogic(r.Context(), svcCtx)
		resp, err := l.GetOnePerCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
