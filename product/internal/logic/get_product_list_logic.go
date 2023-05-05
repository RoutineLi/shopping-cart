package logic

import (
	"context"
	"graduate_design/product/rpc/types/product"
	"strconv"

	"graduate_design/product/internal/svc"
	"graduate_design/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

	for _, id := range pids {
		in := &product.DetailRequest{Id: id}
		out, _ := l.svcCtx.Product.Detail(context.Background(), in)
		item := &types.Product{
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
		resp.Data = append(resp.Data, item)
	}

	resp.Status = "success"
	resp.Code = 200
	resp.Message = "获取全部商品成功, count = " + strconv.Itoa(len(resp.Data))
	return resp, nil
}
