package consume_mqtt

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/define"
	"graduate_design/pkg"
	"graduate_design/websocket/internal/config"
	"graduate_design/websocket/internal/svc"
	"sync"
)

const (
	chanCount   = 10
	bufferCount = 1024
)

type Service struct {
	c        config.Config
	waiter   sync.WaitGroup
	msgsChan []chan *define.KData
}

func NewService(c config.Config) *Service {
	s := &Service{
		c:        c,
		msgsChan: make([]chan *define.KData, chanCount),
	}
	for i := 0; i < chanCount; i++ {
		ch := make(chan *define.KData, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}
	return s
}

func (s *Service) consume(ch chan *define.KData) {
	defer s.waiter.Done()

	for {
		m, ok := <-ch
		if !ok {
			logx.Error("[CONSUME_MQTT ERROR]: consume_mqtt service exit")
		}
		if m != nil {
			define.ClientMap.Range(func(k, v any) bool {
				var cli = k.(*define.Client)
				if cli != nil && v.(string) == m.DeviceKey {
					if cli.Conn != nil {
						cli.Conn.WriteMessage(1, []byte(m.Payload))
					}
					return true
				}
				return false
			})
		}
	}
}

func NewConsumer(ctx *svc.ServiceContext) {
	s := NewService(ctx.Config)
	ctx.Kq = kq.MustNewQueue(ctx.Config.Kafka, kq.WithHandle(s.Consume))
	defer ctx.Kq.Stop()
	ctx.Kq.Start()
}

func (s *Service) Consume(_, value string) error {
	logx.Info("consume value: ", value)
	var data []*define.KData
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		logx.Error(err)
		return err
	}
	for _, d := range data {
		hashVal := pkg.UuidToHash(d.DeviceKey)
		s.msgsChan[hashVal%chanCount] <- d
	}
	return nil
}
