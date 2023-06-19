package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"graduate_design/product/rpc/types/product"

	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequest) (resp *types.ProductListResponse, err error) {
	resp = new(types.ProductListResponse)
	count := 0
	var pids []uint32
	rpcRsp, _ := l.svcCtx.RpcProduct.Ids(context.Background(), &product.IdsRequest{})
	if len(rpcRsp.Ids) == 0 {
		return resp, nil
	}
	for i := int((req.Page - 1) * req.Size); i < len(rpcRsp.Ids); i++ {
		if count == int(req.Size) {
			break
		}
		pids = append(pids, rpcRsp.Ids[i])
		count++
	}
	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pids {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer[*types.Product], cancel func(error)) {
		id := item.(uint32)
		out, err := l.svcCtx.RpcProduct.Detail(l.ctx, &product.DetailRequest{Id: id})
		if err != nil {
			return
		}
		p := &types.Product{
			Name:          out.Data.Name,
			Img:           out.Data.Img,
			Price:         out.Data.Price,
			Origin:        out.Data.Origin,
			Brand:         out.Data.Brand,
			Specification: out.Data.Specification,
			ShelfLife:     out.Data.ShelfLife,
			Description:   out.Data.Description,
			Count:         int(out.Data.Count),
			Type:          out.Data.Type,
			Latitude:      out.Data.Latitude,
			Longitude:     out.Data.Longitude,
			Location:      out.Data.Location,
		}
		writer.Write(p)
	}, func(pipe <-chan *types.Product, writer mr.Writer[[]*types.Product], cancel func(error)) {
		var r []*types.Product
		for p := range pipe {
			r = append(r, p)
		}
		writer.Write(r)
	})

	if err != nil {
		return nil, err
	}
	resp.Count = int64(count)
	resp.List = ps
	return resp, nil
}
