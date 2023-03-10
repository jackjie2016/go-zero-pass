// Code generated by goctl. DO NOT EDIT.
// Source: route.proto

package server

import (
	"context"

	"go-zero-pass/app/service/k8s/route/rpc/internal/logic"
	"go-zero-pass/app/service/k8s/route/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/route/rpc/pb"
)

type RouteServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRouteServer
}

func NewRouteServer(svcCtx *svc.ServiceContext) *RouteServer {
	return &RouteServer{
		svcCtx: svcCtx,
	}
}

// 对外提供添加服务
func (s *RouteServer) AddRoute(ctx context.Context, in *pb.RouteInfo) (*pb.Response, error) {
	l := logic.NewAddRouteLogic(ctx, s.svcCtx)
	return l.AddRoute(in)
}

func (s *RouteServer) DeleteRoute(ctx context.Context, in *pb.RouteId) (*pb.Response, error) {
	l := logic.NewDeleteRouteLogic(ctx, s.svcCtx)
	return l.DeleteRoute(in)
}

func (s *RouteServer) UpdateRoute(ctx context.Context, in *pb.RouteInfo) (*pb.Response, error) {
	l := logic.NewUpdateRouteLogic(ctx, s.svcCtx)
	return l.UpdateRoute(in)
}

func (s *RouteServer) FindRouteByID(ctx context.Context, in *pb.RouteId) (*pb.RouteInfo, error) {
	l := logic.NewFindRouteByIDLogic(ctx, s.svcCtx)
	return l.FindRouteByID(in)
}

func (s *RouteServer) FindAllRoute(ctx context.Context, in *pb.FindAll) (*pb.AllRoute, error) {
	l := logic.NewFindAllRouteLogic(ctx, s.svcCtx)
	return l.FindAllRoute(in)
}
