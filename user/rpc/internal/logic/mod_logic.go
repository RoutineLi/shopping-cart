package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/models"
	"graduate_design/user/rpc/types/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/user/rpc/internal/svc"
)

type ModLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModLogic {
	return &ModLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModLogic) Mod(in *user.ModRequest) (*user.ModResponse, error) {
	err := l.svcCtx.DB.Debug().Where("id = ?", in.Id).Updates(&models.User{
		Nickname: in.Data.Username,
		Password: in.Data.Password,
		Avatar:   in.Data.Avatar,
		Motto:    in.Data.Motto,
		Gender:   in.Data.Gender,
		Age:      uint(in.Data.Age),
		Phone:    in.Data.Phone,
		Email:    in.Data.Email,
	}).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		return &user.ModResponse{Status: false}, err
	}
	threading.GoSafe(func() {
		l.svcCtx.RedisClient.Del(strconv.Itoa(int(in.Id)) + "U")
		return
	})
	return &user.ModResponse{Status: true}, nil
}
