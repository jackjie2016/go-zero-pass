package logic

import (
	"context"
	"encoding/json"
	errorx "go-zero-pass/app/common/error"
	"go-zero-pass/app/service/k8s/route/api/internal/svc"
	"go-zero-pass/app/service/k8s/route/api/internal/types"
	"go-zero-pass/app/service/k8s/route/rpc/route"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRouteByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindRouteByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRouteByIdLogic {
	return &FindRouteByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRouteByIdLogic) FindRouteById(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	logx.Info("Received routeApi.FindRouteById request")

	//获取route信息
	routeInfo, err := l.svcCtx.RouteRpc.FindRouteByID(l.ctx, &route.RouteId{
		Id: req.Id,
	})
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewDefaultError("获取route信息失败")
	}
	//返回route结果
	resp.StatusCode = 200
	b, _ := json.Marshal(routeInfo)
	resp.Body = string(b)
	return resp, nil
}
