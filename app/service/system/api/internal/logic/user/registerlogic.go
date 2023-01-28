package user

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"
	"net/http"

	"go-zero-pass/app/common/message"
	"go-zero-pass/app/service/system/api/internal/logic/captcha"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.SimpleMsg, err error) {
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.CreateOrUpdateUser(context.Background(),
			&core.CreateOrUpdateUserReq{
				Id:       0,
				Username: req.Username,
				Password: req.Password,
				Email:    req.Email,
			})
		if err != nil {
			l.Logger.Error("register logic: create user err: ", err.Error())
			return nil, err
		}
		resp = &types.SimpleMsg{
			Msg: user.Msg,
		}
		return resp, nil
	} else {
		return nil, new_errorx.NewApiError(http.StatusBadRequest, message.WrongCaptcha)
	}
}
