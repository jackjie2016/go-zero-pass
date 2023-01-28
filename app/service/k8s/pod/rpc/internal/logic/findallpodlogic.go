package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/pod/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"
	"go-zero-pass/app/service/k8s/pod/rpc/pod"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllPodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllPodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllPodLogic {
	return &FindAllPodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllPodLogic) FindAllPod(in *pb.FindAll) (*pb.AllPod, error) {
	//查询所有pod
	allPod, err := l.svcCtx.ModelService.FindAllPod()
	if err != nil {
		logx.Error(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	var rsp *pb.AllPod
	//整理格式
	for _, v := range allPod {
		podInfo := &pod.PodInfo{}
		err := common.SwapTo(v, podInfo)
		if err != nil {
			logx.Error(err)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		rsp.PodInfo = append(rsp.PodInfo, podInfo)
	}
	return rsp, nil
}
