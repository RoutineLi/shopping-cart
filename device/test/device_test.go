package test

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"graduate_design/define"
	"graduate_design/device/deviceclient"
	"graduate_design/device/internal/config"
	"graduate_design/device/types/device"
	"graduate_design/models"
	"log"
	"testing"
)

var deviceClient device.DeviceClient

var cFile = flag.String("f", "../etc/device.yaml", "the config file")

func init() {
	var c config.Config
	conf.MustLoad(*cFile, &c)
	deviceClient = deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC))
}

// hb
func client() {
	temp := &models.Product{}
	models.NewDB(define.ProDB)
	var i int
	fmt.Printf("请输入物品id: \n")
	_, _ = fmt.Scanf("%d", i)
	models.DB.Where("id = ?", 4).First(temp)
	msg, _ := json.Marshal(temp)
	reply, err := deviceClient.SendMessage(context.Background(), &device.SendMessageRequest{
		DeviceKey: "77012a0d-4197-42a4-ba13-d6b183a51711",
		Msg:       string(msg),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", reply)
}

// lx
func client_2() {
	models.NewDB(define.ProDB)
	for i := 4; i < 9; i++ {
		temp := &models.Product{}
		models.DB.Where("id = ?", i).First(temp)
		msg, _ := json.Marshal(temp)
		deviceClient.SendMessage(context.Background(), &device.SendMessageRequest{
			DeviceKey: "c0bb490b-0a81-4f7a-9008-9bcb93fa90f0",
			Msg:       string(msg),
		})
	}
}

func TestMessage(t *testing.T) {
	//client_2()
	client()
}
