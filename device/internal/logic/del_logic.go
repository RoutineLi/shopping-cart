package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
	emqx_api "graduate_design/emqx-api"
	"graduate_design/models"
	"strconv"

	"graduate_design/device/internal/svc"
	"graduate_design/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogic) Del(in *device.DelRequest) (*device.DelResponse, error) {
	data := new(models.Device)
	err := l.svcCtx.DB.Model(new(models.Device)).Where("id = ?", in.Id).
		Find(data).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return &device.DelResponse{Status: false}, err
	}
	err = l.svcCtx.DB.Model(new(models.Device)).Transaction(func(tx *gorm.DB) error {
		err = tx.Where("id = ?", in.Id).Delete(new(models.Device)).Error
		if err != nil {
			logx.Error("[DB ERROR]: ", err)
			return err
		}

		//EMQX
		err = emqx_api.DeleteAuthUser(data.Key)
		if err != nil {
			logx.Error("[EMQX ERROR]: ", err)
			return err
		}

		//Redis cache-out
		threading.GoSafe(func() {
			l.svcCtx.RedisClient.Del(strconv.Itoa(int(data.Id)) + "D")
			l.svcCtx.RedisClient.Del(strconv.Itoa(int(data.UserId)) + "BY_USERID")
		})
		return nil
	})
	return &device.DelResponse{Status: true}, nil
}
