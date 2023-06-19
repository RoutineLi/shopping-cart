package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"
	"graduate_design/product/rpc/types/product"
	"strconv"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListLogic) GetProductList(req *types.GetProductListRequest) (resp *types.GetProductListResponse, err error) {
	resp = new(types.GetProductListResponse)
	rsp := new(product.IdsResponse)
	var pids []uint32
	rsp, err = l.svcCtx.Product.Ids(context.Background(), &product.IdsRequest{})
	if err != nil {
		logx.Error("[RPC ERROR]: ", err)
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
	}, func(item interface{}, writer mr.Writer[*types.Product], cancel func(error)) {
		id := item.(uint32)
		out, err := l.svcCtx.Product.Detail(l.ctx, &product.DetailRequest{Id: id})
		if err != nil {
			return
		}
		p := &types.Product{
			Id:            uint(out.Data.Id),
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

	resp.Status = "success"
	resp.Code = 200
	resp.Data = ps
	resp.Message = "获取全部商品成功, count = " + strconv.Itoa(len(resp.Data))

	return resp, nil
}
