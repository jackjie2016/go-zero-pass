package logic

import (
	"context"

	"go-zero-pass/app/service/k8s/volume/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/volume/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVolumeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVolumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVolumeLogic {
	return &DeleteVolumeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteVolumeLogic) DeleteVolume(in *pb.VolumeId) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *volume.DeleteVolume request")
	volumModel, err := l.svcCtx.ModelService.FindVolumeByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//从k8s中删除，并且删除数据库
	if err := l.svcCtx.ModelService.DeleteVolumeFromK8s(volumModel); err != nil {
		logx.Error(err)
		return nil, err
	}
	return &pb.Response{}, nil
}
