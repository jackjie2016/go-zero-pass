package logic

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/message"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/internal/util"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *core.ChangePasswordReq) (*core.BaseResp, error) {
	var target model.User
	result := l.svcCtx.DB.Where("uuid = ?", in.Uuid).First(&target)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, new_errorx.DatabaseError)
	}
	if result.RowsAffected == 0 {
		logx.Errorw("User does not exist", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.NotFound, new_errorx.TargetNotExist)
	}

	if ok := util.BcryptCheck(in.OldPassword, target.Password); ok {
		target.Password = util.BcryptEncrypt(in.NewPassword)
		result = l.svcCtx.DB.Updates(&target)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, new_errorx.DatabaseError)
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
		}
	} else {
		logx.Errorw("Old password is wrong", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.InvalidArgument, message.WrongPassword)
	}
	logx.Infow("Change password successful", logx.Field("UUID", in.Uuid))
	return &core.BaseResp{Msg: new_errorx.UpdateSuccess}, nil
}
