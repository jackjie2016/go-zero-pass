package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/volume/rpc/volume"

	"go-zero-pass/app/service/k8s/volume/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/volume/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllVolumeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllVolumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllVolumeLogic {
	return &FindAllVolumeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllVolumeLogic) FindAllVolume(in *pb.FindAll) (*pb.AllVolume, error) {
	// todo: add your logic here and delete this line
	allVolume, err := l.svcCtx.ModelService.FindAllVolume()
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	rsp := &pb.AllVolume{}
	//整理格式l.svcCtx.ModelService
	for _, v := range allVolume {
		//创建实例
		volumeInfo := &volume.VolumeInfo{}
		//数据转化
		if err := common.SwapTo(v, volumeInfo); err != nil {
			logx.Error(err)
			return nil, err
		}
		//数据合并
		rsp.VolumeInfo = append(rsp.VolumeInfo, volumeInfo)
	}
	return rsp, nil
}
