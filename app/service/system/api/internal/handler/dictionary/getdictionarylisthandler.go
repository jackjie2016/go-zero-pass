package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/dictionary"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /dict/list dictionary getDictionaryList
// Get dictionary list | 获取字典列表
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryListReq
// Responses:
//   200: DictionaryListResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetDictionaryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewGetDictionaryListLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
