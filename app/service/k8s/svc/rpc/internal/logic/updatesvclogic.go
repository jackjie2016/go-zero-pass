package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"

	"go-zero-pass/app/service/k8s/svc/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSvcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSvcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSvcLogic {
	return &UpdateSvcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
	UpdateSvc 对外提供添加服务

测试数据

	{
	  "svcNamespace": "default",
	  "svcName": "gva-web2",
	  "svcPodName": "gva-web",
	  "svcType": "NodePort",
	  "svcTeamId": "1",
	  "svcPort": [
	    {
	      "svcPort": 8080,
	      "svcTargetPort": 8080,
	      "svcNodePort": 30012,
	      "svcPortProtocol": "TCP"
	    }
	  ]
	}
*/
func (l *UpdateSvcLogic) UpdateSvc(in *pb.SvcInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *svc.UpdateSvc request")
	//先更新k8s里面的数据
	if err := l.svcCtx.ModelService.UpdateSvcToK8s(in); err != nil {
		logx.Error(err)
		return nil, err
	}
	//查询数据库中的svc
	service, err := l.svcCtx.ModelService.FindSvcByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//数据类型转换
	if err := common.SwapTo(in, service); err != nil {
		logx.Error(err)
		return nil, err
	}
	//更新到数据中
	if err := l.svcCtx.ModelService.UpdateSvc(service); err != nil {
		logx.Error(err)
		return nil, err
	}
	return &pb.Response{Msg: "ok"}, nil
}
