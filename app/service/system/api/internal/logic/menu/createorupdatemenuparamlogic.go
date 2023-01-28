package menu

import (
	"context"
	"go-zero-pass/app/service/system/rpc/types/core"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuParamLogic {
	return &CreateOrUpdateMenuParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateMenuParamLogic) CreateOrUpdateMenuParam(req *types.CreateOrUpdateMenuParamReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateMenuParam(l.ctx, &core.CreateOrUpdateMenuParamReq{
		Id:     uint64(req.ID),
		MenuId: uint64(req.MenuId),
		Type:   req.DataType,
		Key:    req.Key,
		Value:  req.Value,
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
