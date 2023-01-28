package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/role"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /role/list role getRoleList
// Get role list | 获取角色列表
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PageInfo
// Responses:
//   200: RoleListResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewGetRoleListLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
