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

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *core.IDReq) (*core.BaseResp, error) {
	var children []*model.Menu
	check := l.svcCtx.DB.Where("parent_id = ?", in.ID).Find(&children)
	if check.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", check.Error.Error()))
		return nil, status.Error(codes.Internal, check.Error.Error())
	}
	if check.RowsAffected != 0 {
		logx.Errorw("Delete menu failed, please check its children had been deleted",
			logx.Field("MenuId", in.ID))
		return nil, status.Error(codes.InvalidArgument, message.ChildrenExistError)
	}

	result := l.svcCtx.DB.Delete(&model.Menu{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Delete menu failed, please check the menu id", logx.Field("MenuId", in.ID))
		return nil, status.Error(codes.InvalidArgument, message.MenuNotExists)
	}

	logx.Infow("Delete menu successfully", logx.Field("menuId", in.ID))
	return &core.BaseResp{Msg: new_errorx.DeleteSuccess}, nil
}
