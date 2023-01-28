package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-pass/app/service/system/api/internal/logic/oauth"
	"go-zero-pass/app/service/system/api/internal/svc"
)

// swagger:route POST /oauth/login/callback oauth oauthCallback
// Oauth log in callback route | Oauth 登录返回调接口
// Responses:
//   200: CallbackResp
//   401: SimpleMsg
//   500: SimpleMsg

func OauthCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewOauthCallbackLogic(r, svcCtx)
		resp, err := l.OauthCallback()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
