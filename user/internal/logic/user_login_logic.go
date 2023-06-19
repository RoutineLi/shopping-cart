package logic

import (
	"context"
	"errors"
	"graduate_design/models"
	"graduate_design/pkg"

	"graduate_design/user/internal/svc"
	"graduate_design/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func GenErrorResponse(resp *types.UserLoginResponse, err error) {
	resp.Code = 400
	resp.Message = err.Error()
	resp.Status = "failure"
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	resp = new(types.UserLoginResponse)
	claim := new(models.User)
	err = l.svcCtx.DB.Where("nickname = ? AND password = ?", req.Nickname, pkg.Md5(req.Password)).First(claim).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		err = errors.New("用户名密码错误")
		GenErrorResponse(resp, err)
		return resp, err
	}
	token, err := pkg.GenJwtToken(claim.IsAdmin, claim.Password, claim.Nickname, 3600*24*30, claim.Id)
	if err != nil {
		logx.Error("[DB ERROR]: ", err)
		err = errors.New("token error")
		GenErrorResponse(resp, err)
		return resp, err
	}

	respData := types.UserData{
		Id:       claim.Id,
		Avatar:   claim.Avatar,
		Nickname: claim.Nickname,
		Motto:    claim.Motto,
		Gender:   claim.Gender,
		Age:      claim.Age,
		Phone:    claim.Phone,
		Email:    claim.Email,
		Password: claim.Password,
		Token:    token,
	}
	resp.Data = respData
	resp.Code = 200
	resp.Status = "success"
	resp.Message = "login successful"
	return resp, nil
}
