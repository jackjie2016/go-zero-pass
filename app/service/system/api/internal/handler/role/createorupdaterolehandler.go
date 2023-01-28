package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/role"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /role role createOrUpdateRole
// Create or update role information | 创建或更新角色
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: RoleInfo
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func CreateOrUpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewCreateOrUpdateRoleLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateRole(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
