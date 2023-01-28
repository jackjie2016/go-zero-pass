package logic

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SetRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleStatusLogic {
	return &SetRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRoleStatusLogic) SetRoleStatus(in *core.SetStatusReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Table("roles").Where("id = ?", in.Id).Update("status", in.Status)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Update role status failed, please check the role id", logx.Field("RoleId", in.Id))
		return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
	}

	logx.Infow("Update role status successfully", logx.Field("RoleId", in.Id),
		logx.Field("Status", in.Status))
	return &core.BaseResp{Msg: new_errorx.UpdateSuccess}, nil
}
