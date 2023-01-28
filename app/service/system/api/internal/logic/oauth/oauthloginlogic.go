package oauth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"
)

type OauthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthLoginLogic) OauthLogin(req *types.OauthLoginReq) (resp *types.RedirectResp, err error) {
	result, err := l.svcCtx.CoreRpc.OauthLogin(context.Background(), &core.OauthLoginReq{
		State:    req.State,
		Provider: req.Provider,
	})
	if err != nil {
		return nil, err
	}

	return &types.RedirectResp{URL: result.Url}, nil
}
