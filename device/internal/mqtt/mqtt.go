package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/define"
	"graduate_design/device/internal/svc"
	"graduate_design/models"
	"graduate_design/pkg"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	topic = "/sys/#"
	MC    mqtt.Client
	Ctx   *svc.ServiceContext
	b     *pkg.Batcher
)

const (
	batcherSize     = 100
	batcherBuffer   = 1024
	batcherWorker   = 10
	batcherInterval = time.Second
)

func NewMqttServer(mqttBroker, ClientID, Password string, ctx *svc.ServiceContext) {
	opt := mqtt.NewClientOptions().AddBroker(mqttBroker).
		SetClientID(ClientID).SetUsername("get").SetPassword(Password)

	// 回调
	opt.SetDefaultPublishHandler(publishHandler)

	Ctx = ctx
	MC = mqtt.NewClient(opt)

	// 连接
	if token := MC.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := MC.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer func() {
		// 取消订阅
		if token := MC.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			log.Println("[ERROR] : ", token.Error())
		}
		// 关闭连接
		MC.Disconnect(250)
	}()

	b = pkg.New(
		pkg.WithSize(batcherSize),
		pkg.WithBuffer(batcherBuffer),
		pkg.WithWorker(batcherWorker),
		pkg.WithInterval(batcherInterval),
	)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % batcherWorker
	}

	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*define.KData
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*define.KData))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Error("[Batcher ERROR]: ", err)
		}
		if err = Ctx.KafkaPusher.Push(string(kd)); err != nil {
			logx.Error("[KafkaPusher ERROR]: ", err)
		}
	}
	b.Start()

	select {}
}

// Topic:/sys/key/ping
func publishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("MESSAGE : %s\n", message.Payload())
	fmt.Printf("TOPIC : %s\n", message.Topic())

	topicArray := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")
	if len(topicArray) >= 3 {
		if topicArray[2] == "ping" {
			err := models.UpdateDeviceOnlineTime(topicArray[1])
			if err != nil {
				logx.Error("[DB ERROR]: ", err)
			}
		}
	}

	dkey := topicArray[1]
	hval := pkg.UuidToHash(dkey)
	msg := message.Payload()

	logx.Info("HVAL = ", hval%10)

	if err := b.Add(strconv.FormatInt(hval, 10), &define.KData{
		DeviceKey: dkey,
		Payload:   string(msg),
	}); err != nil {
		logx.Error("[Batcher ERROR]", err)
	}

	logx.Debug("debug")

}
