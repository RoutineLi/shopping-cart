package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/collection"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"graduate_design/define"
	"graduate_design/pkg"
	"graduate_design/product/rpc/types/product"
	"strconv"
	"time"

	"graduate_design/buy/rpc/internal/svc"
	"graduate_design/buy/rpc/types/buy"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	localCacheExpire = time.Second * 60

	batcherSize     = 100
	batcherBuffer   = 100
	batcherWorker   = 10
	batcherInterval = time.Second
)

type SetOrderLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	localCache *collection.Cache
	batcher    *pkg.Batcher
	logx.Logger
}

func NewSetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	s := &SetOrderLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		localCache: localCache,
	}
	b := pkg.New(
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
		var msgs []*define.KBData
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*define.KBData))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Error("[Batcher ERROR]: ", err)
		}
		if err = s.svcCtx.KafkaPusher.Push(string(kd)); err != nil {
			logx.Error("[KafkaPusher ERROR]: ", err)
		}
	}
	s.batcher = b
	s.batcher.Start()
	return s
}

func (l *SetOrderLogic) SetOrder(in *buy.SetOrderRequest) (*buy.SetOrderResponse, error) {
	p, err := l.svcCtx.ProductRPC.Detail(l.ctx, &product.DetailRequest{Id: uint32(in.ProductId)})
	if err != nil {
		return nil, err
	}
	if p.Data.Count <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient count")
	}
	if err = l.batcher.Add(strconv.FormatInt(in.ProductId, 10), &define.KBData{Uid: in.UserId, Pid: in.ProductId}); err != nil {
		logx.Error("[Batcher ERROR]", err)
	}
	return &buy.SetOrderResponse{}, nil
}
