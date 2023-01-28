// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package reclaimspace

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

// ReclaimSpaceControllerClient is the client API for ReclaimSpaceController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReclaimSpaceControllerClient interface {
	// ControllerReclaimSpace is a procedure that gets called on the CSI
	// Controller.
	ControllerReclaimSpace(ctx context.Context, in *ControllerReclaimSpaceRequest, opts ...grpc.CallOption) (*ControllerReclaimSpaceResponse, error)
}

type reclaimSpaceControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewReclaimSpaceControllerClient(cc grpc.ClientConnInterface) ReclaimSpaceControllerClient {
	return &reclaimSpaceControllerClient{cc}
}

func (c *reclaimSpaceControllerClient) ControllerReclaimSpace(ctx context.Context, in *ControllerReclaimSpaceRequest, opts ...grpc.CallOption) (*ControllerReclaimSpaceResponse, error) {
	out := new(ControllerReclaimSpaceResponse)
	err := c.cc.Invoke(ctx, "/reclaimspace.ReclaimSpaceController/ControllerReclaimSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReclaimSpaceControllerServer is the server API for ReclaimSpaceController service.
// All implementations must embed UnimplementedReclaimSpaceControllerServer
// for forward compatibility
type ReclaimSpaceControllerServer interface {
	// ControllerReclaimSpace is a procedure that gets called on the CSI
	// Controller.
	ControllerReclaimSpace(context.Context, *ControllerReclaimSpaceRequest) (*ControllerReclaimSpaceResponse, error)
	mustEmbedUnimplementedReclaimSpaceControllerServer()
}

// UnimplementedReclaimSpaceControllerServer must be embedded to have forward compatible implementations.
type UnimplementedReclaimSpaceControllerServer struct {
}

func (UnimplementedReclaimSpaceControllerServer) ControllerReclaimSpace(context.Context, *ControllerReclaimSpaceRequest) (*ControllerReclaimSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControllerReclaimSpace not implemented")
}
func (UnimplementedReclaimSpaceControllerServer) mustEmbedUnimplementedReclaimSpaceControllerServer() {
}

// UnsafeReclaimSpaceControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReclaimSpaceControllerServer will
// result in compilation errors.
type UnsafeReclaimSpaceControllerServer interface {
	mustEmbedUnimplementedReclaimSpaceControllerServer()
}

func RegisterReclaimSpaceControllerServer(s grpc.ServiceRegistrar, srv ReclaimSpaceControllerServer) {
	s.RegisterService(&ReclaimSpaceController_ServiceDesc, srv)
}

func _ReclaimSpaceController_ControllerReclaimSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControllerReclaimSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReclaimSpaceControllerServer).ControllerReclaimSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reclaimspace.ReclaimSpaceController/ControllerReclaimSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReclaimSpaceControllerServer).ControllerReclaimSpace(ctx, req.(*ControllerReclaimSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReclaimSpaceController_ServiceDesc is the grpc.ServiceDesc for ReclaimSpaceController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReclaimSpaceController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reclaimspace.ReclaimSpaceController",
	HandlerType: (*ReclaimSpaceControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ControllerReclaimSpace",
			Handler:    _ReclaimSpaceController_ControllerReclaimSpace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reclaimspace/reclaimspace.proto",
}

// ReclaimSpaceNodeClient is the client API for ReclaimSpaceNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReclaimSpaceNodeClient interface {
	// NodeReclaimSpace is a procedure that gets called on the CSI NodePlugin.
	NodeReclaimSpace(ctx context.Context, in *NodeReclaimSpaceRequest, opts ...grpc.CallOption) (*NodeReclaimSpaceResponse, error)
}

type reclaimSpaceNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewReclaimSpaceNodeClient(cc grpc.ClientConnInterface) ReclaimSpaceNodeClient {
	return &reclaimSpaceNodeClient{cc}
}

func (c *reclaimSpaceNodeClient) NodeReclaimSpace(ctx context.Context, in *NodeReclaimSpaceRequest, opts ...grpc.CallOption) (*NodeReclaimSpaceResponse, error) {
	out := new(NodeReclaimSpaceResponse)
	err := c.cc.Invoke(ctx, "/reclaimspace.ReclaimSpaceNode/NodeReclaimSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReclaimSpaceNodeServer is the server API for ReclaimSpaceNode service.
// All implementations must embed UnimplementedReclaimSpaceNodeServer
// for forward compatibility
type ReclaimSpaceNodeServer interface {
	// NodeReclaimSpace is a procedure that gets called on the CSI NodePlugin.
	NodeReclaimSpace(context.Context, *NodeReclaimSpaceRequest) (*NodeReclaimSpaceResponse, error)
	mustEmbedUnimplementedReclaimSpaceNodeServer()
}

// UnimplementedReclaimSpaceNodeServer must be embedded to have forward compatible implementations.
type UnimplementedReclaimSpaceNodeServer struct {
}

func (UnimplementedReclaimSpaceNodeServer) NodeReclaimSpace(context.Context, *NodeReclaimSpaceRequest) (*NodeReclaimSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeReclaimSpace not implemented")
}
func (UnimplementedReclaimSpaceNodeServer) mustEmbedUnimplementedReclaimSpaceNodeServer() {}

// UnsafeReclaimSpaceNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReclaimSpaceNodeServer will
// result in compilation errors.
type UnsafeReclaimSpaceNodeServer interface {
	mustEmbedUnimplementedReclaimSpaceNodeServer()
}

func RegisterReclaimSpaceNodeServer(s grpc.ServiceRegistrar, srv ReclaimSpaceNodeServer) {
	s.RegisterService(&ReclaimSpaceNode_ServiceDesc, srv)
}

func _ReclaimSpaceNode_NodeReclaimSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeReclaimSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReclaimSpaceNodeServer).NodeReclaimSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reclaimspace.ReclaimSpaceNode/NodeReclaimSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReclaimSpaceNodeServer).NodeReclaimSpace(ctx, req.(*NodeReclaimSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReclaimSpaceNode_ServiceDesc is the grpc.ServiceDesc for ReclaimSpaceNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReclaimSpaceNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reclaimspace.ReclaimSpaceNode",
	HandlerType: (*ReclaimSpaceNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodeReclaimSpace",
			Handler:    _ReclaimSpaceNode_NodeReclaimSpace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reclaimspace/reclaimspace.proto",
}
