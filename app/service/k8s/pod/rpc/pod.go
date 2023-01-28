package main

import (
	"flag"
	"fmt"
	"go-zero-pass/app/service/k8s/pod/rpc/internal/config"
	"go-zero-pass/app/service/k8s/pod/rpc/internal/server"
	"go-zero-pass/app/service/k8s/pod/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/pod/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/service/k8s/pod/rpc/etc/pod.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPodServer(grpcServer, server.NewPodServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server  at %s...\n", c.ListenOn)
	s.Start()
}
