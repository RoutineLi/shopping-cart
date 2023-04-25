package logic

import (
	"context"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/define"
	"graduate_design/models"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteRequest) (resp *types.ProductDeleteResponse, err error) {
	resp = new(types.ProductDeleteResponse)
	err = l.svcCtx.DB.Model(new(models.Product)).Where("id = ?", req.Id).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Code = 400
		resp.Msg = err.Error()
		return resp, err
	}
	err = l.svcCtx.DB.Debug().Where("id = ?", req.Id).Delete(new(models.Product)).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Code = 400
		resp.Msg = err.Error()
		return resp, err
	}

	//Redis cache-out
	_, err = l.svcCtx.RedisClient.Hdel(define.ProCache, strconv.Itoa(int(req.Id)))
	if err != nil {
		logx.Error("[CACHE ERROR]: ", err)
		resp.Code = 400
		resp.Msg = err.Error()
		return resp, err
	}

	resp.Code = 200
	resp.Msg = "success"
	return resp, nil
}
