package logic

import (
	"context"
	"errors"
	"graduate_design/pkg"
	"graduate_design/user/rpc/types/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/user/rpc/internal/svc"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *user.UserAuthRequest) (*user.UserAuthResponse, error) {
	if in.Token == "" {
		return nil, errors.New("token is null")
	}
	userClaim, err := pkg.CheckJwtToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(user.UserAuthResponse)
	resp.Password = userClaim.Password
	resp.IsAdmin = uint64(userClaim.IsAdmin)
	resp.Extend = map[string]string{
		"name":   userClaim.Name,
		"userid": strconv.Itoa(int(userClaim.Id)),
	}
	return resp, nil
}
