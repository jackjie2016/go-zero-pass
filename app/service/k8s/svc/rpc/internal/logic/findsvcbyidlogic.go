package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"

	"go-zero-pass/app/service/k8s/svc/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSvcByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSvcByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSvcByIDLogic {
	return &FindSvcByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindSvcByIDLogic) FindSvcByID(in *pb.SvcId) (*pb.SvcInfo, error) {
	// todo: add your logic here and delete this line
	logx.Info("查找服务")
	svcModel, err := l.svcCtx.ModelService.FindSvcByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	rsp := &pb.SvcInfo{}
	if err := common.SwapTo(svcModel, rsp); err != nil {
		logx.Error(err)
		return nil, err
	}
	return rsp, nil
}
