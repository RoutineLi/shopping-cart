package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"go/scanner"
	"graduate_design/device/deviceclient"
	"graduate_design/device/internal/config"
	"graduate_design/device/types/device"
	"graduate_design/product/rpc/productclient"
	"graduate_design/product/rpc/types/product"
	"io"
	"log"
	"os"
)

var productClient productclient.Product
var deviceClient device.DeviceClient
var sc scanner.Scanner
var cwd, _ = os.Getwd()
var cFile = flag.String("f", cwd+"/device/etc/device.yaml", "the config file")

type data struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Location  string  `json:"location"`
	Id        int     `json:"id"`
}

type item struct {
	Data   data `json:"data"`
	Status int  `json:"status"`
}

func init() {
	var c config.Config
	conf.MustLoad(*cFile, &c)
	deviceClient = deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC))
	productClient = productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC))
	logx.DisableStat()
}

func main() {
	var i int
	var op int
	for {
		fmt.Printf("请输入物品id和操作: \n")
		_, err := fmt.Scan(&i, &op)
		if err == io.EOF {
			break
		}
		detail, err := productClient.Detail(context.Background(), &product.DetailRequest{Id: uint32(i)})
		if err != nil {
			log.Fatal(err)
		}
		Item := &item{
			Data: data{
				Latitude:  detail.Data.Latitude,
				Longitude: detail.Data.Longitude,
				Location:  detail.Data.Location,
				Id:        int(detail.Data.Id),
			},
			Status: op,
		}
		msg, _ := json.Marshal(Item)
		reply, err := deviceClient.SendMessage(context.Background(), &device.SendMessageRequest{
			DeviceKey: "77012a0d-4197-42a4-ba13-d6b183a51711",
			Msg:       string(msg),
		})
		reply, err = deviceClient.SendMessage(context.Background(), &device.SendMessageRequest{
			DeviceKey: "c0bb490b-0a81-4f7a-9008-9bcb93fa90f0",
			Msg:       string(msg),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", reply)
	}
}
