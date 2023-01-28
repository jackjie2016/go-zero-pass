package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/route/rpc/route"

	"go-zero-pass/app/service/k8s/route/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/route/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllRouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllRouteLogic {
	return &FindAllRouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllRouteLogic) FindAllRoute(in *pb.FindAll) (*pb.AllRoute, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *route.FindAllRoute request")

	rsp := &pb.AllRoute{}
	allRoute, err := l.svcCtx.ModelService.FindAllRoute()

	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//整理下格式
	for _, v := range allRoute {
		//创建实例
		routeInfo := &route.RouteInfo{}
		//把查询出来的数据进行转化
		if err := common.SwapTo(v, routeInfo); err != nil {
			logx.Error(err)
			return nil, err
		}
		//数据合并
		rsp.RouteInfo = append(rsp.RouteInfo, routeInfo)
	}
	return rsp, nil
}
