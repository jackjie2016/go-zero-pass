// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-pass/app/service/k8s/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users/id/:id",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/create",
				Handler: CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/greet/from/:id",
				Handler: GreetHandler(serverCtx),
			},
		},
	)
}