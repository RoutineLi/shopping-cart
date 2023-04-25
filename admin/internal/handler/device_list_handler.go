package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"graduate_design/admin/internal/logic"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
)

func DeviceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		var page int
		var size int
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		size, _ = strconv.Atoi(r.URL.Query().Get("size"))
		req.Page = uint(page)
		req.Size = uint(size)
		req.Name = r.URL.Query().Get("name")

		l := logic.NewDeviceListLogic(r.Context(), svcCtx)
		resp, err := l.DeviceList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
