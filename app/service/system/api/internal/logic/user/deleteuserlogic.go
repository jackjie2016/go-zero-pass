package user

import (
	"context"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.IDReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.DeleteUser(context.Background(), &core.IDReq{ID: uint64(req.ID)})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{
		Msg: data.Msg,
	}, nil
}
