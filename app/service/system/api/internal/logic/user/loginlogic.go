package user

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"
	"net/http"
	"time"

	"go-zero-pass/app/common/message"
	"go-zero-pass/app/service/system/api/internal/logic/captcha"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.Login(context.Background(),
			&core.LoginReq{
				Username: req.Username,
				Password: req.Password,
			})
		if err != nil {
			return nil, err
		}

		token, err := GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, user.Id, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, int64(user.RoleId))
		if err != nil {
			return nil, err
		}
		resp = &types.LoginResp{
			UserId: user.Id,
			Token:  token,
			Expire: uint64(time.Now().Add(time.Second * 259200).Unix()),
			Role: types.RoleInfoSimple{
				Value:    user.RoleValue,
				RoleName: user.RoleName,
			},
		}
		return resp, nil
	} else {
		return nil, new_errorx.NewApiError(http.StatusBadRequest, message.WrongCaptcha)
	}
}

func GetJwtToken(secretKey, uuid string, iat, seconds, roleId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = uuid
	claims["roleId"] = roleId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
