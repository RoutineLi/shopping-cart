// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"graduate_design/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/device/list",
				Handler: DeviceListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/device/create",
				Handler: DeviceAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/admin/device/modify",
				Handler: DeviceModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/admin/device/delete",
				Handler: DeviceDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/product/list",
				Handler: ProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/product/create",
				Handler: ProductAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/admin/product/modify",
				Handler: ProductModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/admin/product/delete",
				Handler: ProductDeleteHandler(serverCtx),
			},
		},
	)
}
