package logic

import (
	"context"
	"gorm.io/gorm"
	"graduate_design/define"
	emqx_api "graduate_design/emqx-api"
	"graduate_design/models"

	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteRequest) (resp *types.DeviceDeleteResponse, err error) {
	resp = new(types.DeviceDeleteResponse)
	device := new(models.Device)
	err = l.svcCtx.DB.Model(new(models.Device)).Where("id = ?", req.Id).
		Find(device).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		resp.Msg = err.Error()
		resp.Code = 400
		return
	}
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Where("id = ?", req.Id).Delete(new(models.Device)).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			resp.Msg = err.Error()
			resp.Code = 400
			return err
		}

		//EMQX
		err = emqx_api.DeleteAuthUser(device.Key)
		if err != nil {
			logx.Error("[EMQX ERROR]: ", err)
			return err
		}
		//Redis cache-out
		_, err = l.svcCtx.RedisClient.Hdel(define.DevCache, device.Key)
		if err != nil {
			logx.Error("[CACHE ERROR]: ", err)
			return err
		}
		return nil
	})

	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
