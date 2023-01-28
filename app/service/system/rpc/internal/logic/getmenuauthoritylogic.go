package logic

import (
	"context"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// authorization management service
func (l *GetMenuAuthorityLogic) GetMenuAuthority(in *core.IDReq) (*core.RoleMenuAuthorityResp, error) {
	var r model.Role
	result := l.svcCtx.DB.Preload("Menus").Where(&model.Role{
		Model: gorm.Model{ID: uint(in.ID)},
	}).First(&r)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	var menuIds []uint64
	for _, v := range r.Menus {
		if v.ID != 1 {
			menuIds = append(menuIds, uint64(v.ID))
		}
	}
	return &core.RoleMenuAuthorityResp{MenuId: menuIds}, nil
}
