package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"
	"graduate_design/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryProductListLogic {
	return &CategoryProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryProductListLogic) CategoryProductList(req *types.CategroyProductListRequest) (resp *types.CategroyProductListResponse, err error) {
	resp = new(types.CategroyProductListResponse)
	rsp := new(product.IdsResponse)
	var pids []uint32
	rsp, err = l.svcCtx.Product.Ids(context.Background(), &product.IdsRequest{})
	if err != nil {
		resp.Status = "failure"
		resp.Code = 400
		resp.Message = err.Error()
		return resp, err
	}
	pids = rsp.Ids

	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pids {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer[*types.ProductBasic], cancel func(error)) {
		id := item.(uint32)
		out, err := l.svcCtx.Product.Detail(l.ctx, &product.DetailRequest{Id: id})
		if err != nil {
			return
		}
		p := &types.ProductBasic{
			Id:    uint(out.Data.Id),
			Name:  out.Data.Name,
			Image: out.Data.Img,
		}
		if out.Data.Type == req.Name {
			writer.Write(p)
		}
	}, func(pipe <-chan *types.ProductBasic, writer mr.Writer[[]*types.ProductBasic], cancel func(error)) {
		var r []*types.ProductBasic
		for p := range pipe {
			r = append(r, p)
		}
		writer.Write(r)
	})

	resp.Data = ps
	resp.Message = "获取" + req.Name + "的商品成功"
	resp.Status = "success"
	resp.Code = 200
	return resp, nil
}
