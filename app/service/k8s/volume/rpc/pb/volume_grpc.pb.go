// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: pb/volume.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VolumeClient is the client API for Volume service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VolumeClient interface {
	//对外提供添加服务
	AddVolume(ctx context.Context, in *VolumeInfo, opts ...grpc.CallOption) (*Response, error)
	DeleteVolume(ctx context.Context, in *VolumeId, opts ...grpc.CallOption) (*Response, error)
	UpdateVolume(ctx context.Context, in *VolumeInfo, opts ...grpc.CallOption) (*Response, error)
	FindVolumeByID(ctx context.Context, in *VolumeId, opts ...grpc.CallOption) (*VolumeInfo, error)
	FindAllVolume(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*AllVolume, error)
}

type volumeClient struct {
	cc grpc.ClientConnInterface
}

func NewVolumeClient(cc grpc.ClientConnInterface) VolumeClient {
	return &volumeClient{cc}
}

func (c *volumeClient) AddVolume(ctx context.Context, in *VolumeInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/volume.Volume/AddVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) DeleteVolume(ctx context.Context, in *VolumeId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/volume.Volume/DeleteVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) UpdateVolume(ctx context.Context, in *VolumeInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/volume.Volume/UpdateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) FindVolumeByID(ctx context.Context, in *VolumeId, opts ...grpc.CallOption) (*VolumeInfo, error) {
	out := new(VolumeInfo)
	err := c.cc.Invoke(ctx, "/volume.Volume/FindVolumeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) FindAllVolume(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*AllVolume, error) {
	out := new(AllVolume)
	err := c.cc.Invoke(ctx, "/volume.Volume/FindAllVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumeServer is the server API for Volume service.
// All implementations must embed UnimplementedVolumeServer
// for forward compatibility
type VolumeServer interface {
	//对外提供添加服务
	AddVolume(context.Context, *VolumeInfo) (*Response, error)
	DeleteVolume(context.Context, *VolumeId) (*Response, error)
	UpdateVolume(context.Context, *VolumeInfo) (*Response, error)
	FindVolumeByID(context.Context, *VolumeId) (*VolumeInfo, error)
	FindAllVolume(context.Context, *FindAll) (*AllVolume, error)
	mustEmbedUnimplementedVolumeServer()
}

// UnimplementedVolumeServer must be embedded to have forward compatible implementations.
type UnimplementedVolumeServer struct {
}

func (UnimplementedVolumeServer) AddVolume(context.Context, *VolumeInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVolume not implemented")
}
func (UnimplementedVolumeServer) DeleteVolume(context.Context, *VolumeId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVolume not implemented")
}
func (UnimplementedVolumeServer) UpdateVolume(context.Context, *VolumeInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVolume not implemented")
}
func (UnimplementedVolumeServer) FindVolumeByID(context.Context, *VolumeId) (*VolumeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVolumeByID not implemented")
}
func (UnimplementedVolumeServer) FindAllVolume(context.Context, *FindAll) (*AllVolume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllVolume not implemented")
}
func (UnimplementedVolumeServer) mustEmbedUnimplementedVolumeServer() {}

// UnsafeVolumeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VolumeServer will
// result in compilation errors.
type UnsafeVolumeServer interface {
	mustEmbedUnimplementedVolumeServer()
}

func RegisterVolumeServer(s grpc.ServiceRegistrar, srv VolumeServer) {
	s.RegisterService(&Volume_ServiceDesc, srv)
}

func _Volume_AddVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).AddVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume.Volume/AddVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).AddVolume(ctx, req.(*VolumeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume.Volume/DeleteVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).DeleteVolume(ctx, req.(*VolumeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_UpdateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).UpdateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume.Volume/UpdateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).UpdateVolume(ctx, req.(*VolumeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_FindVolumeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).FindVolumeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume.Volume/FindVolumeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).FindVolumeByID(ctx, req.(*VolumeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_FindAllVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).FindAllVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume.Volume/FindAllVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).FindAllVolume(ctx, req.(*FindAll))
	}
	return interceptor(ctx, in, info, handler)
}

// Volume_ServiceDesc is the grpc.ServiceDesc for Volume service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Volume_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "volume.Volume",
	HandlerType: (*VolumeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddVolume",
			Handler:    _Volume_AddVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _Volume_DeleteVolume_Handler,
		},
		{
			MethodName: "UpdateVolume",
			Handler:    _Volume_UpdateVolume_Handler,
		},
		{
			MethodName: "FindVolumeByID",
			Handler:    _Volume_FindVolumeByID_Handler,
		},
		{
			MethodName: "FindAllVolume",
			Handler:    _Volume_FindAllVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/volume.proto",
}
