// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"suffix_x/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users/id/:userId",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/create",
				Handler: CreateUserHandler(serverCtx),
			},
		},
	)
}
