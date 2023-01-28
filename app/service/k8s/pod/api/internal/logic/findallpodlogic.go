package logic

import (
	"context"

	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindallpodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindallpodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindallpodLogic {
	return &FindallpodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindallpodLogic) Findallpod(req *types.FindAll) (resp *types.PodList, err error) {
	// todo: add your logic here and delete this line

	return
}
