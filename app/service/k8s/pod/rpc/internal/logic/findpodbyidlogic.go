package logic

import (
	"context"
	customUtils "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/pod/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPodByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPodByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPodByIDLogic {
	return &FindPodByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindPodByIDLogic) FindPodByID(in *pb.PodId) (*pb.PodInfo, error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	//查询pod数据
	podModel, err := l.svcCtx.ModelService.FindPodByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	podinfo := &pb.PodInfo{}
	err = customUtils.SwapTo(podModel, podinfo)
	if err != nil {
		logx.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return podinfo, nil
}
