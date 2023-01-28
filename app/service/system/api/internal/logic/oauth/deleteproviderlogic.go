package oauth

import (
	"context"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProviderLogic {
	return &DeleteProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProviderLogic) DeleteProvider(req *types.IDReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.DeleteProvider(context.Background(), &core.IDReq{
		ID: uint64(req.ID),
	})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}
