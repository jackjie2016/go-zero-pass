package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/k8s/route/api/internal/logic"
	"go-zero-pass/app/service/k8s/route/api/internal/svc"
	"go-zero-pass/app/service/k8s/route/api/internal/types"
)

func UpdateRouteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateRouteLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRoute(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
