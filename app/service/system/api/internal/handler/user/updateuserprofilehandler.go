package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/user"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

// swagger:route POST /user/profile user updateUserProfile
// Update user's profile | 更新用户个人信息
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ProfileReq
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func UpdateUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUpdateUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserProfile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
