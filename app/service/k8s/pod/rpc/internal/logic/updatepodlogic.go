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

type UpdatePodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePodLogic {
	return &UpdatePodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
UpdatePod 请求参考，k8s中podname和podnamespace为搜索的前置条件

	{
	  "podNamespace": "default",
	  "podName": "gva-web2",
	  "podTeamId": "4",
	  "podCpuMax": 30,
	  "podReplicas": 3,
	  "podMemoryMax": 30,
	  "podPort": [
	    {
	      "containerPort": 8080,
	      "protocol": "tcp"
	    }
	  ],
	  "podEnv": [],
	  "podPullPolicy": "IfNotPresent",
	  "podRestart": "Always",
	  "podType": "Rolling",
	  "podImage": "registry.cn-hangzhou.aliyuncs.com/gva/web:latest",
	  "id": "8"
	}
*/
func (l *UpdatePodLogic) UpdatePod(in *pb.PodInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	//查询数据库中的pod
	_, err := l.svcCtx.ModelService.FindPodByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	rsp := &pb.Response{}
	newpodModel := &model.Pod{}
	if err := common.SwapTo(in, newpodModel); err != nil {
		logx.Error(err)

		return &pb.Response{}, err
	}

	//var podModel model.Pod
	if err := l.svcCtx.ModelService.UpdateToK8s(in); err != nil {
		logx.Error(err)

		return &pb.Response{}, err
	} else {
		//操作数据库写入数据
		err := l.svcCtx.ModelService.UpdatePod(newpodModel)
		if err != nil {
			logx.Error(err)
			return &pb.Response{}, err
		}
		logx.Info("Pod 更新成功数据库ID号为：" + strconv.FormatInt(newpodModel.ID, 10))
		rsp.Msg = "Pod 更新成功数据库ID号为：" + strconv.FormatInt(newpodModel.ID, 10)
	}

	return rsp, nil
}
