package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-pass/app/service/k8s/route/api/internal/config"
	"go-zero-pass/app/service/k8s/route/rpc/route"
)

type ServiceContext struct {
	Config   config.Config
	RouteRpc route.Route
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		RouteRpc: route.NewRoute(zrpc.MustNewClient(c.RouteRpc)),
	}
}
