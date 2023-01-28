package logic

import (
	"context"
	"fmt"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"

	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePodLogic {
	return &UpdatePodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePodLogic) UpdatePod(req *types.PodInfo) error {
	// todo: add your logic here and delete this line
	fmt.Println("接受到 podApi.AddPod 的请求")
	resPodInfo := &pb.PodInfo{}
	common.SwapTo(req, resPodInfo)
	if _, err := l.svcCtx.PodRpc.UpdatePod(l.ctx, resPodInfo); err != nil {
		return err
	}
	return nil
}
