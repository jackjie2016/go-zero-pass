// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-pass/app/service/k8s/pod/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/FindPodById/Id/:Id",
				Handler: FindPodByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/FindAllPod",
				Handler: FindallpodHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/AddPod",
				Handler: AddPodHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/DeletePodById/Id/:Id",
				Handler: DeletePodByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/UpdatePod",
				Handler: UpdatePodHandler(serverCtx),
			},
		},
	)
}
