package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/device/deviceclient"
	"graduate_design/device/internal/config"
	"graduate_design/device/types/device"
	"testing"
)

var deviceClient device.DeviceClient

func init() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	deviceClient = deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC))
}

func TestSendMessage(t *testing.T) {
	reply, err := deviceClient.SendMessage(context.Background(), &device.SendMessageRequest{
		DeviceKey: "fdasfds",
		Msg:       "hello world",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", reply)
}
