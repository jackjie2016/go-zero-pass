package logic

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"
	"time"

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

type CreateOrUpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuLogic {
	return &CreateOrUpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuLogic) CreateOrUpdateMenu(in *core.CreateOrUpdateMenuReq) (*core.BaseResp, error) {
	// get parent level
	var menuLevel uint32
	if in.ParentId != 0 {
		var parent model.Menu
		result := l.svcCtx.DB.Where("id = ?", in.ParentId).First(&parent)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("Wrong parent ID", logx.Field("parentId", in.ParentId))
			return nil, status.Error(codes.InvalidArgument, message.ParentNotExist)
		}
		menuLevel = parent.MenuLevel + 1
	} else {
		in.ParentId = 1
		menuLevel = 1
	}
	var data *model.Menu
	if in.Id == 0 {
		// create menu
		data = &model.Menu{
			Model:     gorm.Model{},
			MenuType:  in.MenuType,
			MenuLevel: menuLevel,
			ParentId:  uint(in.ParentId),
			Path:      in.Path,
			Name:      in.Name,
			Redirect:  in.Redirect,
			Component: in.Component,
			OrderNo:   in.OrderNo,
			Disabled:  in.Disabled,
			Meta: model.Meta{
				Title:              in.Meta.Title,
				Icon:               in.Meta.Icon,
				HideMenu:           in.Meta.HideMenu,
				HideBreadcrumb:     in.Meta.HideBreadcrumb,
				CurrentActiveMenu:  in.Meta.CurrentActiveMenu,
				IgnoreKeepAlive:    in.Meta.IgnoreKeepAlive,
				HideTab:            in.Meta.HideTab,
				FrameSrc:           in.Meta.FrameSrc,
				CarryParam:         in.Meta.CarryParam,
				HideChildrenInMenu: in.Meta.HideChildrenInMenu,
				Affix:              in.Meta.Affix,
				DynamicLevel:       in.Meta.DynamicLevel,
				RealPath:           in.Meta.RealPath,
			},
		}
		result := l.svcCtx.DB.Create(data)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, new_errorx.DatabaseError)
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, message.MenuAlreadyExists)
		}

		logx.Infow("Create menu successfully", logx.Field("menuDetail", data))
		return &core.BaseResp{Msg: new_errorx.CreateSuccess}, nil
	} else {
		var origin *model.Menu
		result := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, new_errorx.DatabaseError)
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, message.MenuNotExists)
		}
		data = &model.Menu{
			Model:     gorm.Model{ID: uint(in.Id), CreatedAt: origin.CreatedAt, UpdatedAt: time.Now()},
			MenuLevel: menuLevel,
			MenuType:  in.MenuType,
			ParentId:  uint(in.ParentId),
			Path:      in.Path,
			Name:      in.Name,
			Redirect:  in.Redirect,
			Component: in.Component,
			OrderNo:   in.OrderNo,
			Disabled:  in.Disabled,
			Meta: model.Meta{
				Title:              in.Meta.Title,
				Icon:               in.Meta.Icon,
				HideMenu:           in.Meta.HideMenu,
				HideBreadcrumb:     in.Meta.HideBreadcrumb,
				CurrentActiveMenu:  in.Meta.CurrentActiveMenu,
				IgnoreKeepAlive:    in.Meta.IgnoreKeepAlive,
				HideTab:            in.Meta.HideTab,
				FrameSrc:           in.Meta.FrameSrc,
				CarryParam:         in.Meta.CarryParam,
				HideChildrenInMenu: in.Meta.HideChildrenInMenu,
				Affix:              in.Meta.Affix,
				DynamicLevel:       in.Meta.DynamicLevel,
				RealPath:           in.Meta.RealPath,
			},
		}
		result = l.svcCtx.DB.Save(data)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
		}

		logx.Infow("Update menu successfully", logx.Field("Detail", data))
		return &core.BaseResp{Msg: new_errorx.UpdateSuccess}, nil
	}
}
