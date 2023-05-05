package logic

import (
	"context"
	"graduate_design/user/internal/svc"
	"graduate_design/user/internal/types"
	"graduate_design/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserLogic {
	return &GetAllUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUserLogic) GetAllUser(req *types.GetAllUserRequest) (resp *types.GetAllUserResponse, err error) {
	resp = &types.GetAllUserResponse{}
	rsp := &user.IdsResponse{}
	var uids []uint32
	rsp, err = l.svcCtx.RpcUser.Ids(context.Background(), &user.IdsRequest{})
	if err != nil {
		logx.Error("[RPC ERROR]: ", err)
		return nil, err
	}
	uids = rsp.Ids

	for _, id := range uids {
		in := &user.DetailRequest{Id: id}
		out, _ := l.svcCtx.RpcUser.Detail(context.Background(), in)
		item := &types.UserBasic{
			Id:       uint(out.Id),
			Avatar:   out.Avatar,
			Nickname: out.Username,
			Motto:    out.Motto,
			Gender:   out.Gender,
			Age:      uint(out.Age),
			Phone:    out.Phone,
			Email:    out.Email,
			Password: out.Password,
		}
		resp.Data = append(resp.Data, item)
	}

	resp.Status = "success"
	resp.Code = 200
	resp.Message = "get user all ok"
	return resp, nil
}
