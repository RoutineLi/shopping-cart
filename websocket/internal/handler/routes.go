// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"graduate_design/websocket/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/commodity/shopping_list",
				Handler: GetShoppingListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/service_chat",
				Handler: CustomerServiceChatHandler(serverCtx),
			},
		},
	)
}