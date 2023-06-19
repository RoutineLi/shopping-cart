package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/buy/consume/internal/config"
	"graduate_design/buy/consume/internal/service"
)

var configFile = flag.String("f", "etc/consume.yaml", "the etc file")

func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.Kafka, kq.WithHandle(srv.Consume))
	defer queue.Stop()

	fmt.Println("buy service started!!!")
	queue.Start()
}
