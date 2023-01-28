package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/api"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /api api createOrUpdateApi
// Create or update API information | 创建或更新API
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateApiReq
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func CreateOrUpdateApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewCreateOrUpdateApiLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateApi(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
