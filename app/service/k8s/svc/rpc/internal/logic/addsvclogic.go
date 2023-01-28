package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/svc/model"
	"strconv"

	"go-zero-pass/app/service/k8s/svc/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSvcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSvcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSvcLogic {
	return &AddSvcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
	AddSvc 对外提供添加服务

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
func (l *AddSvcLogic) AddSvc(in *pb.SvcInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	logx.Info("创建服务")
	svcModel := &model.Svc{}
	common.SwapTo(in, svcModel)
	Response := &pb.Response{}
	//到 k8s 中创建服务
	if err := l.svcCtx.ModelService.CreateSvcToK8s(in); err != nil {
		logx.Error(err)
		return nil, err
	} else {
		svcID, err := l.svcCtx.ModelService.AddSvc(svcModel)
		if err != nil {
			//如果逻辑需要自行实现k8s中删除操作
			if err := l.svcCtx.ModelService.DeleteFromK8s(svcModel); err != nil {
				logx.Error(err)
				return nil, err
			}
			logx.Error(err)
			return nil, err
		}
		logx.Info("Svc 添加数据成功ID号为：" + strconv.FormatInt(svcID, 10))
		Response.Msg = "Svc 添加数据成功ID号为：" + strconv.FormatInt(svcID, 10)
	}

	return Response, nil
}
