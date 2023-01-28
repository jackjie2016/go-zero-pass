package menu

import (
	"context"
	"go-zero-pass/app/service/system/rpc/types/core"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuParamLogic {
	return &DeleteMenuParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuParamLogic) DeleteMenuParam(req *types.IDReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.DeleteMenuParam(l.ctx, &core.IDReq{ID: uint64(req.ID)})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
