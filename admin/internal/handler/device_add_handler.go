package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"graduate_design/admin/internal/logic"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
)

func DeviceAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeviceAddLogic(r.Context(), svcCtx)
		resp, err := l.DeviceAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
