package logic

import (
	"context"
	"encoding/json"
	"errors"
	"graduate_design/models"
	"graduate_design/pkg"
	"graduate_design/user/rpc/types/user"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
	"graduate_design/user/rpc/internal/svc"
)

type OpenAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenAuthLogic {
	return &OpenAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenAuthLogic) OpenAuth(in *user.OpenAuthRequest) (*user.OpenAuthResponse, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal(in.Body, &data)
	if err != nil {
		logx.Error("[DB ERROR]: ", err.Error())
		return nil, err
	}

	claim := new(models.User)
	err = l.svcCtx.DB.Model(new(models.User)).Select("app_secret").Where("app_key = ?", data["app_key"]).
		First(claim).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err.Error())
		return nil, err
	}

	arr := make([]string, 0)
	for k, _ := range data {
		arr = append(arr, k)
	}
	sort.Strings(arr)

	var s string
	for _, v := range arr {
		if v != "sign" {
			s += data[v].(string)
		}
	}

	if pkg.Md5(s) != data["sign"].(string) {
		return nil, errors.New("invalid sign")
	}

	return &user.OpenAuthResponse{}, nil
}
