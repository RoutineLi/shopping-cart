package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/models"
	"graduate_design/user/rpc/types/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/user/rpc/internal/svc"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *user.DetailRequest) (*user.DetailResponse, error) {
	item := &models.User{}
	metadata, _ := l.svcCtx.RedisClient.Get(strconv.Itoa(int(in.Id)) + "U")
	if metadata != "" {
		json.Unmarshal([]byte(metadata), item)
		resp := &user.DetailResponse{
			Id:       uint32(item.Id),
			Username: item.Nickname,
			Avatar:   item.Avatar,
			Motto:    item.Motto,
			Gender:   item.Gender,
			Age:      uint32(item.Age),
			Phone:    item.Phone,
			Email:    item.Email,
			Password: item.Password,
		}
		return resp, nil
	}
	err := l.svcCtx.DB.Model(new(models.User)).Where("id = ?", in.Id).First(item).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return nil, err
	}
	resp := &user.DetailResponse{
		Id:       uint32(item.Id),
		Username: item.Nickname,
		Avatar:   item.Avatar,
		Motto:    item.Motto,
		Gender:   item.Gender,
		Age:      uint32(item.Age),
		Phone:    item.Phone,
		Email:    item.Email,
		Password: item.Password,
	}

	threading.GoSafe(func() {
		temp, _ := json.Marshal(resp)
		l.svcCtx.RedisClient.SetnxEx(strconv.Itoa(int(in.Id))+"U", string(temp), 30*60)
	})
	return resp, nil
}
