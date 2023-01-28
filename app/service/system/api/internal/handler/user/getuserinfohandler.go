package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/user"
	"go-zero-pass/app/service/system/api/internal/svc"
)

// swagger:route GET /user/info user getUserInfo
// Get user basic infomation | 获取用户基本信息
// Responses:
//   200: GetUserInfoResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
