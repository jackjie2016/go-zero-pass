package api

import (
	"context"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiLogic {
	return &CreateOrUpdateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateApiLogic) CreateOrUpdateApi(req *types.CreateOrUpdateApiReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateApi(context.Background(),
		&core.ApiInfo{
			Id:          req.Id,
			CreateAt:    req.CreateAt,
			Path:        req.Path,
			Description: req.Description,
			Group:       req.Group,
			Method:      req.Method,
		})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}
