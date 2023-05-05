package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"graduate_design/user/internal/logic"
	"graduate_design/user/internal/svc"
	"graduate_design/user/internal/types"
)

func GetAllUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetAllUserLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
