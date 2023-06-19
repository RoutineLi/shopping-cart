package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"graduate_design/user/internal/svc"
	"graduate_design/user/internal/types"
	"graduate_design/user/rpc/types/user"
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
	us, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, uid := range uids {
			source <- uid
		}
	}, func(item interface{}, writer mr.Writer[*types.UserBasic], cancel func(error)) {
		id := item.(uint32)
		out, err := l.svcCtx.RpcUser.Detail(l.ctx, &user.DetailRequest{Id: id})
		if err != nil {
			return
		}
		u := &types.UserBasic{
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
		writer.Write(u)
	}, func(pipe <-chan *types.UserBasic, writer mr.Writer[[]*types.UserBasic], cancel func(error)) {
		var r []*types.UserBasic
		for u := range pipe {
			r = append(r, u)
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	resp.Data = us
	resp.Status = "success"
	resp.Code = 200
	resp.Message = "get user all ok"
	return resp, nil
}
