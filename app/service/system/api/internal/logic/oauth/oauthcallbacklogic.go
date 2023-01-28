package oauth

import (
	"context"
	"net/http"
	"time"

	"go-zero-pass/app/service/system/api/internal/logic/user"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewOauthCallbackLogic(r *http.Request, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *OauthCallbackLogic) OauthCallback() (resp *types.CallbackResp, err error) {
	result, err := l.svcCtx.CoreRpc.OauthCallback(context.Background(), &core.CallbackReq{
		State: l.r.FormValue("state"),
		Code:  l.r.FormValue("code"),
	})

	if err != nil {
		return nil, err
	}

	token, err := user.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, result.Id, time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, int64(result.RoleId))

	return &types.CallbackResp{
		UserId: result.Id,
		Role: types.RoleInfoSimple{
			RoleName: result.RoleName,
			Value:    result.RoleValue,
		},
		Token:  token,
		Expire: uint64(time.Now().Add(time.Second * 259200).Unix()),
	}, nil
}
