package main

import (
	"flag"
	"fmt"

	"graduate_design/product/rpc/internal/config"
	"graduate_design/product/rpc/internal/server"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.Conf, func(grpcServer *grpc.Server) {
		product.RegisterProductServer(grpcServer, server.NewProductServer(ctx))

		if c.Conf.Mode == service.DevMode || c.Conf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.Conf.ListenOn)
	s.Start()
}
