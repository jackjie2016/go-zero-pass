package authority

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/authority"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /authority/api authority createOrUpdateApiAuthority
// Create or update API authorization information | 创建或更新API权限
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateApiAuthorityReq
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func CreateOrUpdateApiAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateApiAuthorityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authority.NewCreateOrUpdateApiAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateApiAuthority(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
