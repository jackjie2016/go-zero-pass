package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/dictionary"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /dict/detail dictionary createOrUpdateDictionaryDetail
// Create or update dictionary KV information | 创建或更新字典键值信息
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateDictionaryDetailReq
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func CreateOrUpdateDictionaryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateDictionaryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewCreateOrUpdateDictionaryDetailLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateDictionaryDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
