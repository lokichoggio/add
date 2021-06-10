package main

import (
	"flag"
	"fmt"

	"github.com/lokichoggio/add/add"
	"github.com/lokichoggio/add/internal/config"
	"github.com/lokichoggio/add/internal/interceptor"
	"github.com/lokichoggio/add/internal/server"
	"github.com/lokichoggio/add/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/dev/add.yaml", "the config file")
var listenOn = flag.String("listen", "", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	if *listenOn != "" {
		c.ListenOn = *listenOn
	}

	ctx := svc.NewServiceContext(c)
	srv := server.NewAdderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		add.RegisterAdderServer(grpcServer, srv)
	})
	defer s.Stop()

	// 全局拦截器
	s.AddUnaryInterceptors(interceptor.LoggingInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
