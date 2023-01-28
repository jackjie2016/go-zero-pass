package user

import (
	"context"
	"fmt"
	"go-zero-pass/app/common/response/new_errorx"
	"net/http"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermCodeLogic) GetUserPermCode() (resp *types.PermCodeResp, err error) {
	roleId := l.ctx.Value("roleId")
	if roleId == nil {
		return nil, &new_errorx.ApiError{
			Code: http.StatusUnauthorized,
			Msg:  "sys.login.requireLogin",
		}
	}
	return &types.PermCodeResp{Data: []string{fmt.Sprintf("%v", roleId)}}, nil
}
