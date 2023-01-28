package logic

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProfileLogic) UpdateProfile(in *core.UpdateProfileReq) (*core.BaseResp, error) {
	var origin model.User
	result := l.svcCtx.DB.Where("uuid = ?", in.Uuid).First(&origin)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, new_errorx.DatabaseError)
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Fail to find user, please check the UUID", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.NotFound, new_errorx.TargetNotExist)
	}

	origin.Email = in.Email
	origin.Mobile = in.Mobile
	origin.Nickname = in.Nickname
	origin.Avatar = in.Avatar

	result = l.svcCtx.DB.Save(&origin)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, new_errorx.DatabaseError)
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Fail to update the user profile", logx.Field("Detail", origin))
		return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
	}

	return &core.BaseResp{Msg: new_errorx.UpdateSuccess}, nil
}
