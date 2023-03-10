// Code generated by goctl. DO NOT EDIT.
// Source: pod.proto

package pod

import (
	"context"

	"go-zero-pass/app/service/k8s/pod/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AllPod   = pb.AllPod
	FindAll  = pb.FindAll
	PodEnv   = pb.PodEnv
	PodId    = pb.PodId
	PodInfo  = pb.PodInfo
	PodPort  = pb.PodPort
	Request  = pb.Request
	Response = pb.Response

	Pod interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		AddPod(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*Response, error)
		DeletePod(ctx context.Context, in *PodId, opts ...grpc.CallOption) (*Response, error)
		FindPodByID(ctx context.Context, in *PodId, opts ...grpc.CallOption) (*PodInfo, error)
		UpdatePod(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*Response, error)
		FindAllPod(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*AllPod, error)
	}

	defaultPod struct {
		cli zrpc.Client
	}
)

func NewPod(cli zrpc.Client) Pod {
	return &defaultPod{
		cli: cli,
	}
}

func (m *defaultPod) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewPodClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultPod) AddPod(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewPodClient(m.cli.Conn())
	return client.AddPod(ctx, in, opts...)
}

func (m *defaultPod) DeletePod(ctx context.Context, in *PodId, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewPodClient(m.cli.Conn())
	return client.DeletePod(ctx, in, opts...)
}

func (m *defaultPod) FindPodByID(ctx context.Context, in *PodId, opts ...grpc.CallOption) (*PodInfo, error) {
	client := pb.NewPodClient(m.cli.Conn())
	return client.FindPodByID(ctx, in, opts...)
}

func (m *defaultPod) UpdatePod(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewPodClient(m.cli.Conn())
	return client.UpdatePod(ctx, in, opts...)
}

func (m *defaultPod) FindAllPod(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*AllPod, error) {
	client := pb.NewPodClient(m.cli.Conn())
	return client.FindAllPod(ctx, in, opts...)
}
