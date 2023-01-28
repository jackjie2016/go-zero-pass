package oauth

import (
	"context"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProviderListLogic {
	return &GetProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProviderListLogic) GetProviderList(req *types.PageInfo) (resp *types.ProviderListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetProviderList(context.Background(),
		&core.PageInfoReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ProviderListResp{}
	resp.Total = data.GetTotal()
	for _, v := range data.Data {
		resp.Data = append(resp.Data,
			types.ProviderInfo{
				Id:           v.Id,
				CreateAt:     v.CreateAt,
				Name:         v.Name,
				ClientID:     v.ClientId,
				ClientSecret: v.ClientSecret,
				RedirectURL:  v.RedirectUrl,
				Scopes:       v.Scopes,
				AuthURL:      v.AuthUrl,
				TokenURL:     v.TokenUrl,
				InfoURL:      v.InfoUrl,
				AuthStyle:    int(v.AuthStyle),
			})
	}
	return resp, nil
}
