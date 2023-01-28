package user

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"
	"net/http"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	if l.ctx.Value("userId").(string) == "" {
		return nil, new_errorx.NewApiError(http.StatusUnauthorized, "Please log in")
	}
	user, err := l.svcCtx.CoreRpc.GetUserById(context.Background(),
		&core.UUIDReq{UUID: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		UUID:     user.UUID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Roles: types.GetUserRoleInfo{
			RoleName: user.RoleName,
			Value:    user.RoleValue,
		},
	}, nil
}
