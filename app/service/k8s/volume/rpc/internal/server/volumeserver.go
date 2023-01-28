// Code generated by goctl. DO NOT EDIT.
// Source: volume.proto

package server

import (
	"context"

	"go-zero-pass/app/service/k8s/volume/rpc/internal/logic"
	"go-zero-pass/app/service/k8s/volume/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/volume/rpc/pb"
)

type VolumeServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedVolumeServer
}

func NewVolumeServer(svcCtx *svc.ServiceContext) *VolumeServer {
	return &VolumeServer{
		svcCtx: svcCtx,
	}
}

// 对外提供添加服务
func (s *VolumeServer) AddVolume(ctx context.Context, in *pb.VolumeInfo) (*pb.Response, error) {
	l := logic.NewAddVolumeLogic(ctx, s.svcCtx)
	return l.AddVolume(in)
}

func (s *VolumeServer) DeleteVolume(ctx context.Context, in *pb.VolumeId) (*pb.Response, error) {
	l := logic.NewDeleteVolumeLogic(ctx, s.svcCtx)
	return l.DeleteVolume(in)
}

func (s *VolumeServer) UpdateVolume(ctx context.Context, in *pb.VolumeInfo) (*pb.Response, error) {
	l := logic.NewUpdateVolumeLogic(ctx, s.svcCtx)
	return l.UpdateVolume(in)
}

func (s *VolumeServer) FindVolumeByID(ctx context.Context, in *pb.VolumeId) (*pb.VolumeInfo, error) {
	l := logic.NewFindVolumeByIDLogic(ctx, s.svcCtx)
	return l.FindVolumeByID(in)
}

func (s *VolumeServer) FindAllVolume(ctx context.Context, in *pb.FindAll) (*pb.AllVolume, error) {
	l := logic.NewFindAllVolumeLogic(ctx, s.svcCtx)
	return l.FindAllVolume(in)
}
