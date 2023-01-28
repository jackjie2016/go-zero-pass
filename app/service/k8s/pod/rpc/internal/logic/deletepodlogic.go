package logic

import (
	"context"
	"strconv"

	"go-zero-pass/app/service/k8s/pod/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePodLogic {
	return &DeletePodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePodLogic) DeletePod(in *pb.PodId) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	//先查找数据
	podModel, err := l.svcCtx.ModelService.FindPodByID(in.Id)
	rsp := &pb.Response{}
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if err := l.svcCtx.ModelService.DeleteFromK8s(podModel); err != nil {
		logx.Error(err)
		return nil, err
	}
	rsp.Msg = "Pod 添加成功数据库ID号为：" + strconv.FormatInt(in.Id, 10)
	return rsp, nil
}
