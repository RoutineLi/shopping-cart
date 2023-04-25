package logic

import (
	"context"
	"encoding/json"
	"graduate_design/define"
	"graduate_design/models"

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
	flag := true
	var count int64
	var m map[string]string
	m, _ = l.svcCtx.RedisClient.Hgetall(define.ProCache)
	if len(m) != 0 {
		for k := range m {
			slice := []byte(m[k])
			temp := new(types.Product)
			err = json.Unmarshal(slice, temp)
			if err != nil {
				logx.Error("[CACHE ERROR]: ", err)
				resp.Status = "failure"
				resp.Code = 400
				resp.Message = err.Error()
				return
			}
			resp.Data = append(resp.Data, temp)
		}
		resp.Status = "success"
		resp.Code = 200
		resp.Message = "获取全部商品成功, count = " + string(len(m))
		return resp, nil
	} else if len(m) == 0 {
		flag = false
	}

	err = models.GetProductList("").Count(&count).Find(&resp.Data).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Status = "failure"
		resp.Code = 400
		resp.Message = err.Error()
		return resp, err
	}
	if !flag {
		for _, v := range resp.Data {
			data, _ := json.Marshal(v)
			m[string(v.Id)] = string(data)
		}
		l.svcCtx.RedisClient.Hmset(define.ProCache, m)
	}
	resp.Status = "success"
	resp.Code = 200
	resp.Message = "获取全部商品成功, count = " + string(count)
	return resp, nil
}
