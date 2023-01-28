package logic

import (
	"context"

	"go-zero-pass/app/service/k8s/svc/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSvcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSvcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSvcLogic {
	return &DeleteSvcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSvcLogic) DeleteSvc(in *pb.SvcId) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	logx.Info("删除服务")
	service, err := l.svcCtx.ModelService.FindSvcByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	if err := l.svcCtx.ModelService.DeleteFromK8s(service); err != nil {
		logx.Error(err)
		return nil, err
	}

	return &pb.Response{}, nil
}
