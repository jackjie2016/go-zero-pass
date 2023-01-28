package core

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/core"
	"go-zero-pass/app/service/system/api/internal/svc"
)

// swagger:route get /core/init/database core initDatabase
// Initialize database | 初始化数据库
// Responses:
//   200: SimpleMsg
//   500: SimpleMsg

func InitDatabaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewInitDatabaseLogic(r.Context(), svcCtx)
		resp, err := l.InitDatabase()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
