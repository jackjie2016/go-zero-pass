package captcha

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-pass/app/service/system/api/internal/logic/captcha"
	"go-zero-pass/app/service/system/api/internal/svc"
)

// swagger:route GET /captcha captcha getCaptcha
// Get captcha | 获取验证码
// Responses:
//   200: CaptchaInfo
//   401: SimpleMsg
//   500: SimpleMsg

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
