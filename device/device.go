package main

import (
	"flag"
	"fmt"
	"graduate_design/device/internal/mqtt"

	"graduate_design/device/internal/config"
	"graduate_design/device/internal/server"
	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/device.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		device.RegisterDeviceServer(grpcServer, server.NewDeviceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	go mqtt.NewMqttServer(c.Mqtt.Broker, c.Mqtt.ClientID, c.Mqtt.Password)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
