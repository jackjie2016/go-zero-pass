package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/k8s/volume/model"
	"strconv"

	"go-zero-pass/app/service/k8s/volume/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/volume/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVolumeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVolumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVolumeLogic {
	return &AddVolumeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 对外提供添加服务
func (l *AddVolumeLogic) AddVolume(in *pb.VolumeInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	logx.Info("Received *volume.AddVolume request")
	rsp := &pb.Response{}
	volume := &model.Volume{}
	if err := common.SwapTo(in, volume); err != nil {
		logx.Error(err)
		rsp.Msg = err.Error()
		return nil, err
	}
	//创建volume
	if err := l.svcCtx.ModelService.CreateVolumeToK8s(in); err != nil {
		logx.Error(err)
		rsp.Msg = err.Error()
		return nil, err
	} else {
		//写入数据库
		volumeID, err := l.svcCtx.ModelService.AddVolume(volume)
		if err != nil {
			logx.Error(err)
			rsp.Msg = err.Error()
			return nil, err
		}
		rsp.Msg = "volume 添加成功 ID 号为：" + strconv.FormatInt(volumeID, 10)
	}
	return rsp, nil
}
