package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/pod/model"
	"go-zero-pass/app/service/k8s/pod/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPodLogic {
	return &AddPodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
AddPod grpcui 调试数据

	{
	  "podPort": [
	    {
	      "containerPort": 8080,
	      "protocol": "tcp"
	    }
	  ],
	  "podEnv": [],
	  "podNamespace": "default",
	  "podName": "gva-web2",
	  "podTeamId": "1",
	  "podCpuMax": 30,
	  "podReplicas": 1,
	  "podMemoryMax": 30,
	  "podPullPolicy": "IfNotPresent",
	  "podRestart": "Always",
	  "podType": "Rolling",
	  "podImage": "registry.cn-hangzhou.aliyuncs.com/gva/web:latest"
	}
*/
func (l *AddPodLogic) AddPod(in *pb.PodInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	//podModel := &model.Pod{}
	logx.Info("添加pod")
	podModel := &model.Pod{}
	rsp := &pb.Response{}
	if err := common.SwapTo(in, podModel); err != nil {
		logx.Error(err)
		rsp.Msg = err.Error()
		return nil, err
	}
	if err := l.svcCtx.ModelService.CreateToK8s(in); err != nil {
		logx.Error(err)

		return &pb.Response{}, err
	} else {
		//操作数据库写入数据
		podID, err := l.svcCtx.ModelService.AddPod(podModel)
		if err != nil {
			logx.Error(err)
			return &pb.Response{}, err
		}
		logx.Info("Pod 添加成功数据库ID号为：" + strconv.FormatInt(podID, 10))
		rsp.Msg = "Pod 添加成功数据库ID号为：" + strconv.FormatInt(podID, 10)
	}

	return rsp, nil
}
