package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/dictionary"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /dict/detail/list dictionary getDetailByDictionaryName
// Get dictionary detail list by dictionary name | 根据字典名获取字典键值列表
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryDetailReq
// Responses:
//   200: DictionaryDetailListResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetDetailByDictionaryNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewGetDetailByDictionaryNameLogic(r.Context(), svcCtx)
		resp, err := l.GetDetailByDictionaryName(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
