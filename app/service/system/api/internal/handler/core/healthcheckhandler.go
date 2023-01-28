package core

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/core"
	"go-zero-pass/app/service/system/api/internal/svc"
)

// swagger:route get /core/health core healthCheck
// Check the system status | 检查系统状态

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewHealthCheckLogic(r.Context(), svcCtx)
		err := l.HealthCheck()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
