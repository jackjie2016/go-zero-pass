package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/menu"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /menu/param menu createOrUpdateMenuParam
// Create or update menu parameters | 创建或更新菜单参数
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateMenuParamReq
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func CreateOrUpdateMenuParamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateMenuParamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewCreateOrUpdateMenuParamLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateMenuParam(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
