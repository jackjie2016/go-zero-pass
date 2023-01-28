package logic

import (
	"context"
	"strconv"

	"go-zero-pass/app/service/k8s/route/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/route/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRouteLogic {
	return &DeleteRouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRouteLogic) DeleteRoute(in *pb.RouteId) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *route.DeleteRoute request")
	rsp := &pb.Response{}
	routeModel, err := l.svcCtx.ModelService.FindRouteByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//从k8s中删除，并且删除数据库中数据
	if err := l.svcCtx.ModelService.DeleteRouteFromK8s(routeModel); err != nil {
		logx.Error(err)
		return nil, err
	}
	rsp.Msg = "Route 成功删除 ID 号为：" + strconv.FormatInt(in.Id, 10)
	return rsp, nil
}
