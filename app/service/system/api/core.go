//		Simple Admin
//
//		This is simple admin api doc
//
//		Schemes: http, https
//		Host: localhost:9100
//		BasePath: /
//		Version: 0.0.6
//		Contact: yuansu.china.work@gmail.com
//		SecurityDefinitions:
//		  Token:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//		Security:
//		  - Token: []
//	    Consumes:
//		  - application/json
//
//		Produces:
//		  - application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"

	"go-zero-pass/app/service/system/api/internal/config"
	"go-zero-pass/app/service/system/api/internal/handler"
	"go-zero-pass/app/service/system/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/service/system/api/etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
