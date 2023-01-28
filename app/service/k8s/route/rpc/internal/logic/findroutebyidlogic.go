package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"

	"go-zero-pass/app/service/k8s/route/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/route/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRouteByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRouteByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRouteByIDLogic {
	return &FindRouteByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindRouteByIDLogic) FindRouteByID(in *pb.RouteId) (*pb.RouteInfo, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *route.FindRouteByID request")

	rsp := &pb.RouteInfo{}
	routeModel, err := l.svcCtx.ModelService.FindRouteByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//数据转化
	if err := common.SwapTo(routeModel, rsp); err != nil {
		logx.Error(err)
		return nil, err
	}
	return rsp, nil
}
