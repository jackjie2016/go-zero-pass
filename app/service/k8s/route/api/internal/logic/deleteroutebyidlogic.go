package logic

import (
	"context"
	"encoding/json"
	errorx "go-zero-pass/app/common/error"
	"go-zero-pass/app/service/k8s/route/rpc/route"
	"strconv"

	"go-zero-pass/app/service/k8s/route/api/internal/svc"
	"go-zero-pass/app/service/k8s/route/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRouteByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRouteByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRouteByIdLogic {
	return &DeleteRouteByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRouteByIdLogic) DeleteRouteById(req *types.RouteInfo) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	logx.Info("Received routeApi.DeleteRouteById request")
	if _, ok := req.Get["route_id"]; !ok {
		resp.StatusCode = 500
		return nil, errorx.NewDefaultError("参数异常")
	}
	//获取 route id
	routeIdString := req.Get["route_id"].Values[0]
	routeId, err := strconv.ParseInt(routeIdString, 10, 64)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//调用route 删除服务
	response, err := l.svcCtx.RouteRpc.DeleteRoute(l.ctx, &route.RouteId{
		Id: routeId,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	resp.StatusCode = 200
	b, _ := json.Marshal(response)
	resp.Body = string(b)
	return resp, nil
}
