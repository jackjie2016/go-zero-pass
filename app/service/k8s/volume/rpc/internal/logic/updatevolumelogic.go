package logic

import (
	"context"
	"go-zero-pass/app/service/k8s/volume/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/volume/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVolumeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVolumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVolumeLogic {
	return &UpdateVolumeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVolumeLogic) UpdateVolume(in *pb.VolumeInfo) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	logx.Info("Received *volume.AddVolume request")
	//
	//volumeModel, err := l.svcCtx.ModelService.FindVolumeByID(in.Id)
	//if err != nil {
	//	logx.Error(err)
	//	return nil, err
	//}
	//rsp := &pb.Response{}
	//volume := &model.Volume{}
	//if err := common.SwapTo(in, volume); err != nil {
	//	logx.Error(err)
	//	rsp.Msg = err.Error()
	//	return nil, err
	//}
	////更新volume
	//if err := l.svcCtx.ModelService.UpdateVolumeToK8s(in); err != nil {
	//	logx.Error(err)
	//	rsp.Msg = err.Error()
	//	return nil, err
	//} else {
	//	//写入数据库
	//	err := l.svcCtx.ModelService.UpdateVolume(volume)
	//	if err != nil {
	//		logx.Error(err)
	//		rsp.Msg = err.Error()
	//		return nil, err
	//	}
	//	rsp.Msg = "volume 更新 ID 号为：" + strconv.FormatInt(volumeModel.ID, 10)
	//}
	return &pb.Response{}, nil

}
