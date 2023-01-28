package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/oauth"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /oauth/provider oauth createOrUpdateProvider
// Create or update Provider information | 创建或更新提供商
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateProviderReq
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func CreateOrUpdateProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateProviderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewCreateOrUpdateProviderLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateProvider(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
