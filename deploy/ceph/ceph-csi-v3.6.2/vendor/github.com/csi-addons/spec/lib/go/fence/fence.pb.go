// Code generated by make; DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: fence/fence.proto

package fence

import (
	_ "github.com/container-storage-interface/spec/lib/go/csi"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// FenceClusterNetworkRequest contains the information needed to identify
// the storage cluster so that the appropriate fencing operation can be
// performed.
type FenceClusterNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugin specific parameters passed in as opaque key-value pairs.
	Parameters map[string]string `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Secrets required by the plugin to complete the request.
	Secrets map[string]string `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// list of CIDR blocks on which the fencing operation is expected to be
	// performed.
	Cidrs []*CIDR `protobuf:"bytes,3,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
}

func (x *FenceClusterNetworkRequest) Reset() {
	*x = FenceClusterNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fence_fence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FenceClusterNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FenceClusterNetworkRequest) ProtoMessage() {}

func (x *FenceClusterNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fence_fence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FenceClusterNetworkRequest.ProtoReflect.Descriptor instead.
func (*FenceClusterNetworkRequest) Descriptor() ([]byte, []int) {
	return file_fence_fence_proto_rawDescGZIP(), []int{0}
}

func (x *FenceClusterNetworkRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *FenceClusterNetworkRequest) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *FenceClusterNetworkRequest) GetCidrs() []*CIDR {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

// FenceClusterNetworkResponse is returned by the CSI-driver as a result of
// the FenceClusterNetworkRequest call.
type FenceClusterNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FenceClusterNetworkResponse) Reset() {
	*x = FenceClusterNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fence_fence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FenceClusterNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FenceClusterNetworkResponse) ProtoMessage() {}

func (x *FenceClusterNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fence_fence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FenceClusterNetworkResponse.ProtoReflect.Descriptor instead.
func (*FenceClusterNetworkResponse) Descriptor() ([]byte, []int) {
	return file_fence_fence_proto_rawDescGZIP(), []int{1}
}

// UnfenceClusterNetworkRequest contains the information needed to identify
// the cluster so that the appropriate fence operation can be
// disabled.
type UnfenceClusterNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugin specific parameters passed in as opaque key-value pairs.
	Parameters map[string]string `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Secrets required by the plugin to complete the request.
	Secrets map[string]string `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// list of CIDR blocks on which the fencing operation is expected to be
	// performed.
	Cidrs []*CIDR `protobuf:"bytes,3,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
}

func (x *UnfenceClusterNetworkRequest) Reset() {
	*x = UnfenceClusterNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fence_fence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfenceClusterNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfenceClusterNetworkRequest) ProtoMessage() {}

func (x *UnfenceClusterNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fence_fence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfenceClusterNetworkRequest.ProtoReflect.Descriptor instead.
func (*UnfenceClusterNetworkRequest) Descriptor() ([]byte, []int) {
	return file_fence_fence_proto_rawDescGZIP(), []int{2}
}

func (x *UnfenceClusterNetworkRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *UnfenceClusterNetworkRequest) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *UnfenceClusterNetworkRequest) GetCidrs() []*CIDR {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

// UnfenceClusterNetworkResponse is returned by the CSI-driver as a result of
// the UnfenceClusterNetworkRequest call.
type UnfenceClusterNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnfenceClusterNetworkResponse) Reset() {
	*x = UnfenceClusterNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fence_fence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfenceClusterNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfenceClusterNetworkResponse) ProtoMessage() {}

func (x *UnfenceClusterNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fence_fence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfenceClusterNetworkResponse.ProtoReflect.Descriptor instead.
func (*UnfenceClusterNetworkResponse) Descriptor() ([]byte, []int) {
	return file_fence_fence_proto_rawDescGZIP(), []int{3}
}

// ListClusterFenceRequest contains the information needed to identify
// the cluster so that the appropriate fenced clients can be listed.
type ListClusterFenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugin specific parameters passed in as opaque key-value pairs.
	Parameters map[string]string `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Secrets required by the plugin to complete the request.
	Secrets map[string]string `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListClusterFenceRequest) Reset() {
	*x = ListClusterFenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fence_fence_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClusterFenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClusterFenceRequest) ProtoMessage() {}

func (x *ListClusterFenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fence_fence_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClusterFenceRequest.ProtoReflect.Descriptor instead.
func (*ListClusterFenceRequest) Descriptor() ([]byte, []int) {
	return file_fence_fence_proto_rawDescGZIP(), []int{4}
}

func (x *ListClusterFenceRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *ListClusterFenceRequest) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

// ListClusterFenceResponse holds the information about the result of the
// ListClusterFenceResponse call.
type ListClusterFenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of IPs that are blocklisted by the SP.
	Cidrs []*CIDR `protobuf:"bytes,1,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
}

func (x *ListClusterFenceResponse) Reset() {
	*x = ListClusterFenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fence_fence_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClusterFenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClusterFenceResponse) ProtoMessage() {}

func (x *ListClusterFenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fence_fence_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClusterFenceResponse.ProtoReflect.Descriptor instead.
func (*ListClusterFenceResponse) Descriptor() ([]byte, []int) {
	return file_fence_fence_proto_rawDescGZIP(), []int{5}
}

func (x *ListClusterFenceResponse) GetCidrs() []*CIDR {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

// CIDR holds a CIDR block.
type CIDR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CIDR block
	Cidr string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
}

func (x *CIDR) Reset() {
	*x = CIDR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fence_fence_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIDR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIDR) ProtoMessage() {}

func (x *CIDR) ProtoReflect() protoreflect.Message {
	mi := &file_fence_fence_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIDR.ProtoReflect.Descriptor instead.
func (*CIDR) Descriptor() ([]byte, []int) {
	return file_fence_fence_proto_rawDescGZIP(), []int{6}
}

func (x *CIDR) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

var File_fence_fence_proto protoreflect.FileDescriptor

var file_fence_fence_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x73, 0x69, 0x2f, 0x63, 0x73, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc,
	0x02, 0x0a, 0x1a, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x4d, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x03, 0x98, 0x42, 0x01, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12,
	0x21, 0x0a, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x52, 0x05, 0x63, 0x69, 0x64,
	0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1d, 0x0a,
	0x1b, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe2, 0x02, 0x0a,
	0x1c, 0x55, 0x6e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x4f, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0x98, 0x42, 0x01, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x52,
	0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x6e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xb0, 0x02, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e,
	0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4a,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0x98, 0x42,
	0x01, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x52, 0x05, 0x63,
	0x69, 0x64, 0x72, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x43, 0x49, 0x44, 0x52, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72,
	0x32, 0xae, 0x02, 0x0a, 0x0f, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x13, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x21, 0x2e, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x15, 0x55, 0x6e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x23, 0x2e,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1e,
	0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fence_fence_proto_rawDescOnce sync.Once
	file_fence_fence_proto_rawDescData = file_fence_fence_proto_rawDesc
)

func file_fence_fence_proto_rawDescGZIP() []byte {
	file_fence_fence_proto_rawDescOnce.Do(func() {
		file_fence_fence_proto_rawDescData = protoimpl.X.CompressGZIP(file_fence_fence_proto_rawDescData)
	})
	return file_fence_fence_proto_rawDescData
}

var file_fence_fence_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_fence_fence_proto_goTypes = []interface{}{
	(*FenceClusterNetworkRequest)(nil),    // 0: fence.FenceClusterNetworkRequest
	(*FenceClusterNetworkResponse)(nil),   // 1: fence.FenceClusterNetworkResponse
	(*UnfenceClusterNetworkRequest)(nil),  // 2: fence.UnfenceClusterNetworkRequest
	(*UnfenceClusterNetworkResponse)(nil), // 3: fence.UnfenceClusterNetworkResponse
	(*ListClusterFenceRequest)(nil),       // 4: fence.ListClusterFenceRequest
	(*ListClusterFenceResponse)(nil),      // 5: fence.ListClusterFenceResponse
	(*CIDR)(nil),                          // 6: fence.CIDR
	nil,                                   // 7: fence.FenceClusterNetworkRequest.ParametersEntry
	nil,                                   // 8: fence.FenceClusterNetworkRequest.SecretsEntry
	nil,                                   // 9: fence.UnfenceClusterNetworkRequest.ParametersEntry
	nil,                                   // 10: fence.UnfenceClusterNetworkRequest.SecretsEntry
	nil,                                   // 11: fence.ListClusterFenceRequest.ParametersEntry
	nil,                                   // 12: fence.ListClusterFenceRequest.SecretsEntry
}
var file_fence_fence_proto_depIdxs = []int32{
	7,  // 0: fence.FenceClusterNetworkRequest.parameters:type_name -> fence.FenceClusterNetworkRequest.ParametersEntry
	8,  // 1: fence.FenceClusterNetworkRequest.secrets:type_name -> fence.FenceClusterNetworkRequest.SecretsEntry
	6,  // 2: fence.FenceClusterNetworkRequest.cidrs:type_name -> fence.CIDR
	9,  // 3: fence.UnfenceClusterNetworkRequest.parameters:type_name -> fence.UnfenceClusterNetworkRequest.ParametersEntry
	10, // 4: fence.UnfenceClusterNetworkRequest.secrets:type_name -> fence.UnfenceClusterNetworkRequest.SecretsEntry
	6,  // 5: fence.UnfenceClusterNetworkRequest.cidrs:type_name -> fence.CIDR
	11, // 6: fence.ListClusterFenceRequest.parameters:type_name -> fence.ListClusterFenceRequest.ParametersEntry
	12, // 7: fence.ListClusterFenceRequest.secrets:type_name -> fence.ListClusterFenceRequest.SecretsEntry
	6,  // 8: fence.ListClusterFenceResponse.cidrs:type_name -> fence.CIDR
	0,  // 9: fence.FenceController.FenceClusterNetwork:input_type -> fence.FenceClusterNetworkRequest
	2,  // 10: fence.FenceController.UnfenceClusterNetwork:input_type -> fence.UnfenceClusterNetworkRequest
	4,  // 11: fence.FenceController.ListClusterFence:input_type -> fence.ListClusterFenceRequest
	1,  // 12: fence.FenceController.FenceClusterNetwork:output_type -> fence.FenceClusterNetworkResponse
	3,  // 13: fence.FenceController.UnfenceClusterNetwork:output_type -> fence.UnfenceClusterNetworkResponse
	5,  // 14: fence.FenceController.ListClusterFence:output_type -> fence.ListClusterFenceResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_fence_fence_proto_init() }
func file_fence_fence_proto_init() {
	if File_fence_fence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fence_fence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FenceClusterNetworkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fence_fence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FenceClusterNetworkResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fence_fence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfenceClusterNetworkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fence_fence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfenceClusterNetworkResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fence_fence_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClusterFenceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fence_fence_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClusterFenceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fence_fence_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIDR); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fence_fence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fence_fence_proto_goTypes,
		DependencyIndexes: file_fence_fence_proto_depIdxs,
		MessageInfos:      file_fence_fence_proto_msgTypes,
	}.Build()
	File_fence_fence_proto = out.File
	file_fence_fence_proto_rawDesc = nil
	file_fence_fence_proto_goTypes = nil
	file_fence_fence_proto_depIdxs = nil
}
