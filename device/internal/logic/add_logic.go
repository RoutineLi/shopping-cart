package logic

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
	emqx_api "graduate_design/emqx-api"
	"graduate_design/models"
	"graduate_design/pkg"
	"strconv"

	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *device.AddRequest) (*device.AddResponse, error) {
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		device := &models.Device{
			UserId:         uint(in.Userid),
			Name:           in.Name,
			Key:            uuid.New().String(),
			Secret:         uuid.New().String(),
			LastOnlineTime: 0,
		}
		err := tx.Create(device).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return err
		}

		//EMQX认证
		err = emqx_api.CreateAuthUser(&emqx_api.CreateAuthUserRequest{
			UserId:   device.Key,
			Password: pkg.Md5(device.Key + device.Secret),
		})
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return err
		}

		threading.GoSafe(func() {
			data, _ := json.Marshal(device)
			l.svcCtx.RedisClient.Setex(strconv.Itoa(int(device.Id))+"D", string(data), 30*60)
		})
		return nil
	})
	if err != nil {
		return &device.AddResponse{Status: false}, err
	}
	return &device.AddResponse{Status: true}, nil
}
