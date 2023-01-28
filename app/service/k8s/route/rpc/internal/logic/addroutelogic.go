package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/route/model"
	"strconv"

	"go-zero-pass/app/service/k8s/route/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/route/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRouteLogic {
	return &AddRouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
	AddRoute 对外提供添加服务

测试数据

	{
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

*
*/
func (l *AddRouteLogic) AddRoute(in *pb.RouteInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	rsp := &pb.Response{}
	//创建route到k8s
	if err := l.svcCtx.ModelService.CreateRouteToK8s(in); err != nil {
		logx.Error(err)
		rsp.Msg = err.Error()
		return nil, err
	} else {
		//写入数据库
		route := &model.Route{}
		common.SwapTo(in, route)
		routeID, err := l.svcCtx.ModelService.AddRoute(route)
		if err != nil {
			logx.Error(err)
			rsp.Msg = err.Error()
			return nil, err
		}
		logx.Info("Route 添加成功 ID 号为：" + strconv.FormatInt(routeID, 10))
		rsp.Msg = "Route 添加成功 ID 号为：" + strconv.FormatInt(routeID, 10)
	}
	return rsp, nil
}
