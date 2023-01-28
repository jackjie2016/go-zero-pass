package handler

import (
	errorx "go-zero-pass/app/common/error"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/k8s/pod/api/internal/logic"
	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"
)

func DeletePodByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			httpx.Error(w, errorx.HandleGrpcErrorToHttp(err))
			return
		}

		l := logic.NewDeletePodByIdLogic(r.Context(), svcCtx)
		err := l.DeletePodById(&req)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			httpx.Error(w, errorx.HandleGrpcErrorToHttp(err))
		} else {
			httpx.Ok(w)
		}
	}
}
