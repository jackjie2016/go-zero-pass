package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/api"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /api/list api getApiList
// Get API list | 获取API列表
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ApiListReq
// Responses:
//   200: ApiListResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetApiListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApiListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewGetApiListLogic(r.Context(), svcCtx)
		resp, err := l.GetApiList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
