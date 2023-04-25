package logic

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"graduate_design/admin/internal/svc"
	"graduate_design/admin/internal/types"
	"graduate_design/define"
	emqx_api "graduate_design/emqx-api"
	"graduate_design/models"
	"graduate_design/pkg"
)

type DeviceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceAddLogic {
	return &DeviceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceAddLogic) DeviceAdd(req *types.DeviceAddRequest) (resp *types.DeviceAddResponse, err error) {
	resp = new(types.DeviceAddResponse)
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		device := &models.Device{
			UserId:         req.UserId,
			Name:           req.Name,
			Key:            uuid.New().String(),
			Secret:         uuid.New().String(),
			LastOnlineTime: 0,
		}
		err = tx.Create(device).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			resp.Code = 400
			resp.Msg = err.Error()
			return err
		}

		//EMQX认证
		err = emqx_api.CreateAuthUser(&emqx_api.CreateAuthUserRequest{
			UserId:   device.Key,
			Password: pkg.Md5(device.Key + device.Secret),
		})
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			resp.Code = 400
			resp.Msg = err.Error()
			return err
		}

		//Redis cache-in
		data, _ := json.Marshal(device)

		err = l.svcCtx.RedisClient.Hmset(define.DevCache, map[string]string{device.Key: string(data)})

		if err != nil {
			logx.Error("[CACHE ERROR]: ", err)
			resp.Code = 400
			resp.Msg = err.Error()
			return err
		}
		return nil
	})
	if err != nil {
		return resp, err
	}
	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
