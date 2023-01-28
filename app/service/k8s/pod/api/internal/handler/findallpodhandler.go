package handler

import (
	errorx "go-zero-pass/app/common/error"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/k8s/pod/api/internal/logic"
	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"
)

func FindallpodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindAll
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.HandleGrpcErrorToHttp(err))
			return
		}

		l := logic.NewFindallpodLogic(r.Context(), svcCtx)
		resp, err := l.Findallpod(&req)
		if err != nil {
			httpx.Error(w, errorx.HandleGrpcErrorToHttp(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
