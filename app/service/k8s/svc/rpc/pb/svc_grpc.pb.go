// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: pb/svc.proto

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

// SvcClient is the client API for Svc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SvcClient interface {
	//对外提供添加服务
	AddSvc(ctx context.Context, in *SvcInfo, opts ...grpc.CallOption) (*Response, error)
	DeleteSvc(ctx context.Context, in *SvcId, opts ...grpc.CallOption) (*Response, error)
	UpdateSvc(ctx context.Context, in *SvcInfo, opts ...grpc.CallOption) (*Response, error)
	FindSvcByID(ctx context.Context, in *SvcId, opts ...grpc.CallOption) (*SvcInfo, error)
	FindAllSvc(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*AllSvc, error)
}

type svcClient struct {
	cc grpc.ClientConnInterface
}

func NewSvcClient(cc grpc.ClientConnInterface) SvcClient {
	return &svcClient{cc}
}

func (c *svcClient) AddSvc(ctx context.Context, in *SvcInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/svc.Svc/AddSvc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcClient) DeleteSvc(ctx context.Context, in *SvcId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/svc.Svc/DeleteSvc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcClient) UpdateSvc(ctx context.Context, in *SvcInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/svc.Svc/UpdateSvc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcClient) FindSvcByID(ctx context.Context, in *SvcId, opts ...grpc.CallOption) (*SvcInfo, error) {
	out := new(SvcInfo)
	err := c.cc.Invoke(ctx, "/svc.Svc/FindSvcByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcClient) FindAllSvc(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*AllSvc, error) {
	out := new(AllSvc)
	err := c.cc.Invoke(ctx, "/svc.Svc/FindAllSvc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SvcServer is the server API for Svc service.
// All implementations must embed UnimplementedSvcServer
// for forward compatibility
type SvcServer interface {
	//对外提供添加服务
	AddSvc(context.Context, *SvcInfo) (*Response, error)
	DeleteSvc(context.Context, *SvcId) (*Response, error)
	UpdateSvc(context.Context, *SvcInfo) (*Response, error)
	FindSvcByID(context.Context, *SvcId) (*SvcInfo, error)
	FindAllSvc(context.Context, *FindAll) (*AllSvc, error)
	mustEmbedUnimplementedSvcServer()
}

// UnimplementedSvcServer must be embedded to have forward compatible implementations.
type UnimplementedSvcServer struct {
}

func (UnimplementedSvcServer) AddSvc(context.Context, *SvcInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSvc not implemented")
}
func (UnimplementedSvcServer) DeleteSvc(context.Context, *SvcId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSvc not implemented")
}
func (UnimplementedSvcServer) UpdateSvc(context.Context, *SvcInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSvc not implemented")
}
func (UnimplementedSvcServer) FindSvcByID(context.Context, *SvcId) (*SvcInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSvcByID not implemented")
}
func (UnimplementedSvcServer) FindAllSvc(context.Context, *FindAll) (*AllSvc, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllSvc not implemented")
}
func (UnimplementedSvcServer) mustEmbedUnimplementedSvcServer() {}

// UnsafeSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SvcServer will
// result in compilation errors.
type UnsafeSvcServer interface {
	mustEmbedUnimplementedSvcServer()
}

func RegisterSvcServer(s grpc.ServiceRegistrar, srv SvcServer) {
	s.RegisterService(&Svc_ServiceDesc, srv)
}

func _Svc_AddSvc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SvcInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcServer).AddSvc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc.Svc/AddSvc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcServer).AddSvc(ctx, req.(*SvcInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Svc_DeleteSvc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SvcId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcServer).DeleteSvc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc.Svc/DeleteSvc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcServer).DeleteSvc(ctx, req.(*SvcId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Svc_UpdateSvc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SvcInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcServer).UpdateSvc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc.Svc/UpdateSvc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcServer).UpdateSvc(ctx, req.(*SvcInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Svc_FindSvcByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SvcId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcServer).FindSvcByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc.Svc/FindSvcByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcServer).FindSvcByID(ctx, req.(*SvcId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Svc_FindAllSvc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcServer).FindAllSvc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc.Svc/FindAllSvc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcServer).FindAllSvc(ctx, req.(*FindAll))
	}
	return interceptor(ctx, in, info, handler)
}

// Svc_ServiceDesc is the grpc.ServiceDesc for Svc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Svc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.Svc",
	HandlerType: (*SvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSvc",
			Handler:    _Svc_AddSvc_Handler,
		},
		{
			MethodName: "DeleteSvc",
			Handler:    _Svc_DeleteSvc_Handler,
		},
		{
			MethodName: "UpdateSvc",
			Handler:    _Svc_UpdateSvc_Handler,
		},
		{
			MethodName: "FindSvcByID",
			Handler:    _Svc_FindSvcByID_Handler,
		},
		{
			MethodName: "FindAllSvc",
			Handler:    _Svc_FindAllSvc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/svc.proto",
}
