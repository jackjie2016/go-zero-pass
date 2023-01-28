package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"

	"go-zero-pass/app/service/k8s/volume/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/volume/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVolumeByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVolumeByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVolumeByIDLogic {
	return &FindVolumeByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindVolumeByIDLogic) FindVolumeByID(in *pb.VolumeId) (*pb.VolumeInfo, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *volume.FindVolumeByID request")
	volumeModel, err := l.svcCtx.ModelService.FindVolumeByID(in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	rsp := &pb.VolumeInfo{}
	//数据转化
	if err := common.SwapTo(volumeModel, rsp); err != nil {
		logx.Error(err)
		return nil, err
	}

	return rsp, nil
}
