package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/k8s/route/api/internal/logic"
	"go-zero-pass/app/service/k8s/route/api/internal/svc"
	"go-zero-pass/app/service/k8s/route/api/internal/types"
)

func DeleteRouteByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RouteInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteRouteByIdLogic(r.Context(), svcCtx)
		resp, err := l.DeleteRouteById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
