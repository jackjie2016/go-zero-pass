package oauth

import (
	"context"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateProviderLogic {
	return &CreateOrUpdateProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateProviderLogic) CreateOrUpdateProvider(req *types.CreateOrUpdateProviderReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateProvider(context.Background(),
		&core.ProviderInfo{
			Id:           req.Id,
			Name:         req.Name,
			ClientId:     req.ClientID,
			ClientSecret: req.ClientSecret,
			RedirectUrl:  req.RedirectURL,
			Scopes:       req.Scopes,
			AuthUrl:      req.AuthURL,
			TokenUrl:     req.TokenURL,
			AuthStyle:    uint64(req.AuthStyle),
			InfoUrl:      req.InfoURL,
			CreateAt:     req.CreateAt,
		})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}
