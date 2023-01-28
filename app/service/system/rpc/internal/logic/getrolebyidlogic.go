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

type GetRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleByIdLogic) GetRoleById(in *core.IDReq) (*core.RoleInfo, error) {
	var role model.Role
	result := l.svcCtx.DB.Where("id = ?", in.ID).First(&role)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Fail to find the role, please check the role id", logx.Field("RoleId", in.ID))
		return nil, status.Error(codes.InvalidArgument, new_errorx.GetInfoFailed)
	}
	return &core.RoleInfo{
		Id:            uint64(role.ID),
		Name:          role.Name,
		Value:         role.Value,
		DefaultRouter: role.DefaultRouter,
		Status:        role.Status,
		Remark:        role.Remark,
		OrderNo:       role.OrderNo,
		CreateAt:      role.CreatedAt.UnixMilli(),
	}, nil
}
