package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-pass/app/service/k8s/pod/api/internal/config"
	"go-zero-pass/app/service/k8s/pod/rpc/pod"
)

type ServiceContext struct {
	Config config.Config
	PodRpc pod.Pod
}

func NewServiceContext(c config.Config) *ServiceContext {

	PodRpc := pod.NewPod(zrpc.MustNewClient(c.PodRpc))
	if PodRpc == nil {
		logx.Error("链接失败")
		return nil
	}
	logx.Error("PodRpc 链接成功")
	return &ServiceContext{
		Config: c,

		PodRpc: PodRpc,
	}
}
