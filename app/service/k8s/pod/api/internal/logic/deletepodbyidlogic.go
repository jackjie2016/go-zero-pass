package logic

import (
	"context"
	"fmt"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"

	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePodByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePodByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePodByIdLogic {
	return &DeletePodByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePodByIdLogic) DeletePodById(req *types.Request) error {
	// todo: add your logic here and delete this line
	fmt.Println("接受到 podApi.AddPod 的请求")
	if _, err := l.svcCtx.PodRpc.DeletePod(l.ctx, &pb.PodId{Id: req.Id}); err != nil {
		return err
	}
	return nil
}
