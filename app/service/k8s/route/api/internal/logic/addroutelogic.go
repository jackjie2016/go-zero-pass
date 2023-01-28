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

type AddRouteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRouteLogic {
	return &AddRouteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRouteLogic) AddRoute(req *types.RouteInfo) (resp *types.Response, err error) {
	logx.Info("Received routeApi.AddRoute request")
	addRouteInfo := &route.RouteInfo{}
	routePathName, ok := req.Post["route_path_name"]
	if ok && len(routePathName.Values) > 0 {
		port, err := strconv.ParseInt(req.Post["route_backend_service_port"].Values[0], 10, 32)
		if err != nil {
			logx.Error(err)
			return nil, errorx.NewDefaultError("参数异常")
		}
		//这里如果有多个Path需要处理多个
		routePath := &route.RoutePath{
			RoutePathName:           req.Post["route_path_name"].Values[0],
			RouteBackendService:     req.Post["route_backend_service"].Values[0],
			RouteBackendServicePort: int32(port),
		}
		//合并
		addRouteInfo.RoutePath = append(addRouteInfo.RoutePath, routePath)
	}

	response, err := l.svcCtx.RouteRpc.AddRoute(l.ctx, addRouteInfo)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	resp.StatusCode = 200
	b, _ := json.Marshal(response)
	resp.Body = string(b)
	return resp, nil
}
