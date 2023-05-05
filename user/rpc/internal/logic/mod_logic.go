package logic

import (
	"context"

	"graduate_design/user/rpc/internal/svc"
	"graduate_design/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &user.ModResponse{}, nil
}
