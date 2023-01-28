package handler

import (
	"github.com/go-playground/validator/v10"
	errorx "go-zero-pass/app/common/error"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/k8s/pod/api/internal/logic"
	"go-zero-pass/app/service/k8s/pod/api/internal/svc"
	"go-zero-pass/app/service/k8s/pod/api/internal/types"
)

func FindPodByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, errorx.NewDefaultError(err.Error()))
			return
		}
		//增加个翻译器
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
			return
		}

		l := logic.NewFindPodByIdLogic(r.Context(), svcCtx)
		resp, err := l.FindPodById(&req)

		if err != nil {
			httpx.Error(w, errorx.HandleGrpcErrorToHttp(err))
			//httpx.ErrorCtx(r.Context(), w, errorx.NewDefaultError(err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
