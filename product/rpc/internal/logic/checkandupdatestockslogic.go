package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndUpdateStocksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAndUpdateStocksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndUpdateStocksLogic {
	return &CheckAndUpdateStocksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	luaScript = `
		local counts = redis.call("HMGET", KEYS[1], "total", "secbuy")
		local total = tonumber(counts[1])
		local secbuy = tonumber(counts[2])
		if secbuy + 1 <= total then
			redis.call("HINCRBY", KEYS[1], "secbuy", 1)
			return 1
		end
		return 0
    `
)

func (l *CheckAndUpdateStocksLogic) CheckAndUpdateStocks(in *product.CAURequest) (*product.CAUResponse, error) {
	val, err := l.svcCtx.RedisClient.EvalCtx(l.ctx, luaScript, []string{stockKey(int64(in.Pid))})
	if err != nil {
		return nil, err
	}
	if val.(int64) == 0 {
		return nil, status.Errorf(codes.ResourceExhausted, fmt.Sprintf("insufficient stock: %d", in.Pid))
	}
	return &product.CAUResponse{}, nil
}

func stockKey(pid int64) string {
	return fmt.Sprintf("stock:%d", pid)
}
