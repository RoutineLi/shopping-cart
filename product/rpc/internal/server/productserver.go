// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"graduate_design/product/rpc/internal/logic"
	"graduate_design/product/rpc/internal/svc"
	"graduate_design/product/rpc/types/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Add(ctx context.Context, in *product.AddRequest) (*product.AddResponse, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}

func (s *ProductServer) Del(ctx context.Context, in *product.DelRequest) (*product.DelResponse, error) {
	l := logic.NewDelLogic(ctx, s.svcCtx)
	return l.Del(in)
}

func (s *ProductServer) Mod(ctx context.Context, in *product.ModRequest) (*product.ModResponse, error) {
	l := logic.NewModLogic(ctx, s.svcCtx)
	return l.Mod(in)
}

func (s *ProductServer) Detail(ctx context.Context, in *product.DetailRequest) (*product.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *ProductServer) Categories(ctx context.Context, in *product.CategoriesRequest) (*product.CategoriesResponse, error) {
	l := logic.NewCategoriesLogic(ctx, s.svcCtx)
	return l.Categories(in)
}

func (s *ProductServer) Ids(ctx context.Context, in *product.IdsRequest) (*product.IdsResponse, error) {
	l := logic.NewIdsLogic(ctx, s.svcCtx)
	return l.Ids(in)
}