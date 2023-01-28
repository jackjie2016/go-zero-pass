package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"

	"go-zero-pass/app/service/k8s/pod/rpc/pb"

	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPodByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPodByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPodByIdLogic {
	return &FindPodByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPodByIdLogic) FindPodById(req *types.Request) (resp *types.PodInfo, err error) {
	// todo: add your logic here and delete this line

	resPodInfo := &pb.PodInfo{}
	common.SwapTo(req, resPodInfo)
	if podInfo, err := l.svcCtx.PodRpc.FindPodByID(l.ctx, &pb.PodId{Id: req.Id}); err != nil {
 
		return nil, err
	} else {
		Response := &types.PodInfo{}
		common.SwapTo(podInfo, Response)
		return Response, nil
	}

}
