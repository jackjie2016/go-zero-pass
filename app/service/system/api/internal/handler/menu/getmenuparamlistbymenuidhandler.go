package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/menu"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /menu/param/list menu getMenuParamListByMenuId
// Get menu extra parameters by menu ID | 获取某个菜单的额外参数列表
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
// Responses:
//   200: MenuParamListByMenuIdResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetMenuParamListByMenuIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewGetMenuParamListByMenuIdLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuParamListByMenuId(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
