package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/define"
	"graduate_design/models"
	"graduate_design/pkg"
	"strconv"

	"graduate_design/user/internal/svc"
	"graduate_design/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	resp = new(types.UserRegisterResponse)
	claim := new(models.User)
	res := l.svcCtx.DB.Where("nickname = ?", req.Nickname).First(claim)
	if res.RowsAffected == 0 {
		claim = &models.User{
			Nickname: req.Nickname,
			Password: pkg.Md5(req.Password),
			Avatar:   req.Avatar,
			Motto:    req.Motto,
			Gender:   req.Gender,
			Age:      req.Age,
			Phone:    req.Phone,
			Email:    req.Email,
			IsAdmin:  req.IsAdmin,
		}
		l.svcCtx.DB.Create(&claim)
		threading.GoSafe(func() {
			temp, _ := json.Marshal(claim)
			l.svcCtx.RedisClient.Sadd(define.UserIdsCache, strconv.Itoa(int(claim.Id)))
			l.svcCtx.RedisClient.Setex(strconv.Itoa(int(claim.Id))+"U", string(temp), 30*60)
		})
		resp.Code = 200
		resp.Status = "success"
		resp.Message = "Registration successful"

		return resp, nil
	} else {
		resp.Code = 400
		resp.Status = "failure"
		resp.Message = "当前用户名已存在"
		logx.Error("debug")
		return resp, nil
	}
	return
}
