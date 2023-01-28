package logic

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/message"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *core.IDReq) (*core.BaseResp, error) {
	var users []model.User
	check := l.svcCtx.DB.Model(&model.User{}).Where("role_id = ?", in.ID).Find(&users).RowsAffected
	if check != 0 {
		logx.Errorw("Delete role failed, please check the users who belongs to the role had been deleted",
			logx.Field("RoleId", in.ID))
		return nil, status.Error(codes.InvalidArgument, message.UserExists)
	}
	result := l.svcCtx.DB.Delete(&model.Role{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Delete role failed, please check the role id", logx.Field("RoleId", in.ID))
		return nil, status.Error(codes.InvalidArgument, new_errorx.DeleteFailed)
	}

	logx.Infow("Delete role successfully", logx.Field("RoleId", in.ID))
	return &core.BaseResp{Msg: new_errorx.DeleteSuccess}, nil
}
