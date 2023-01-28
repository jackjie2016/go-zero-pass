package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"

	"go-zero-pass/app/service/k8s/route/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/route/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRouteLogic {
	return &UpdateRouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
UpdateRoute
测试数据

	{
	  "routePath": [],
	  "id": "3",
	  "routeName": "gva-web-ingress2",
	  "routeNamespace": "default",
	  "routeHost": "www.code688.com",
	 "routePath": [
		    {
		      "routePathName": "/",
		      "routeBackendService": "gva-web",
		      "routeBackendServicePort": 8888
		    }
		  ]
	}
*/
func (l *UpdateRouteLogic) UpdateRoute(in *pb.RouteInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *route.UpdateRoute request")
	if err := l.svcCtx.ModelService.UpdateRouteToK8s(in); err != nil {
		logx.Error(err)
		return nil, err
	}
	//查询数据库的信息
	routeModel, err := l.svcCtx.ModelService.FindRouteByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//数据更新
	if err := common.SwapTo(in, routeModel); err != nil {
		logx.Error(err)
		return nil, err
	}

	if err := l.svcCtx.ModelService.UpdateRoute(routeModel); err != nil {
		logx.Error(err)
		return nil, err
	}

	return &pb.Response{Msg: "更新成功"}, nil
}
