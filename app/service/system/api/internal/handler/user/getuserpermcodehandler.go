package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/user"
	"go-zero-pass/app/service/system/api/internal/svc"
)

// swagger:route GET /user/perm user getUserPermCode
// Get user's permission code | 获取用户权限码
// Responses:
//   200: PermCodeResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetUserPermCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserPermCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetUserPermCode()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
