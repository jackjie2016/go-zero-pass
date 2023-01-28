package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	errorx "go-zero-pass/app/common/error"
	"go-zero-pass/app/common/initialize"
	"net/http"

	"go-zero-pass/app/service/k8s/pod/api/internal/config"
	"go-zero-pass/app/service/k8s/pod/api/internal/handler"
	"go-zero-pass/app/service/k8s/pod/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/service/k8s/pod/api/etc/pod-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	//引入自定义的表单验证 并且加入翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err.Error())
	}

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		fmt.Println(err)
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
