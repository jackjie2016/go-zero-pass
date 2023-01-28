package dictionary

import (
	"context"
	"go-zero-pass/app/service/system/rpc/types/core"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryDetailLogic {
	return &CreateOrUpdateDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateDictionaryDetailLogic) CreateOrUpdateDictionaryDetail(req *types.CreateOrUpdateDictionaryDetailReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateDictionaryDetail(context.Background(), &core.DictionaryDetail{
		Id:       req.Id,
		Title:    req.Title,
		Key:      req.Key,
		Value:    req.Value,
		Status:   req.Status,
		ParentId: int64(req.ParentId),
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
