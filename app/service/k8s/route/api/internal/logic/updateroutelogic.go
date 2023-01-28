package logic

import (
	"context"
	"encoding/json"

	"go-zero-pass/app/service/k8s/route/api/internal/svc"
	"go-zero-pass/app/service/k8s/route/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRouteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRouteLogic {
	return &UpdateRouteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRouteLogic) UpdateRoute(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	logx.Info("Received routeApi.UpdateRoute request")
	resp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routeApi/UpdateRoute'}")
	resp.Body = string(b)

	return
}
