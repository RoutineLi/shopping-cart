// Code generated by goctl. DO NOT EDIT.
// Source: buy.proto

package buyclient

import (
	"context"

	"graduate_design/buy/rpc/types/buy"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BuyProductsRequest  = buy.BuyProductsRequest
	BuyProductsResponse = buy.BuyProductsResponse
	Product             = buy.Product
	SetOrderRequest     = buy.SetOrderRequest
	SetOrderResponse    = buy.SetOrderResponse

	Buy interface {
		SetProducts(ctx context.Context, in *BuyProductsRequest, opts ...grpc.CallOption) (*BuyProductsResponse, error)
		SetOrder(ctx context.Context, in *SetOrderRequest, opts ...grpc.CallOption) (*SetOrderResponse, error)
	}

	defaultBuy struct {
		cli zrpc.Client
	}
)

func NewBuy(cli zrpc.Client) Buy {
	return &defaultBuy{
		cli: cli,
	}
}

func (m *defaultBuy) SetProducts(ctx context.Context, in *BuyProductsRequest, opts ...grpc.CallOption) (*BuyProductsResponse, error) {
	client := buy.NewBuyClient(m.cli.Conn())
	return client.SetProducts(ctx, in, opts...)
}

func (m *defaultBuy) SetOrder(ctx context.Context, in *SetOrderRequest, opts ...grpc.CallOption) (*SetOrderResponse, error) {
	client := buy.NewBuyClient(m.cli.Conn())
	return client.SetOrder(ctx, in, opts...)
}
