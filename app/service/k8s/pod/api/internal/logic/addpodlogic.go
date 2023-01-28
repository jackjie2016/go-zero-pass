package logic

import (
	"context"
	"fmt"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPodLogic {
	return &AddPodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPodLogic) AddPod(req *types.PodInfo) error {
	fmt.Println("接受到 podApi.AddPod 的请求")
	resPodInfo := &pb.PodInfo{}
	common.SwapTo(req, resPodInfo)
	if _, err := l.svcCtx.PodRpc.AddPod(l.ctx, resPodInfo); err != nil {
		return err
	}
	return nil
}
