package menu

import (
	"context"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

type CreateOrUpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuLogic {
	return &CreateOrUpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateMenuLogic) CreateOrUpdateMenu(req *types.CreateOrUpdateMenuReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateMenu(context.Background(), &core.CreateOrUpdateMenuReq{
		Id:        uint64(req.ID),
		MenuType:  req.MenuType,
		ParentId:  uint32(req.ParentId),
		Path:      req.Path,
		Name:      req.Name,
		Redirect:  req.Redirect,
		Component: req.Component,
		OrderNo:   req.OrderNo,
		Disabled:  req.Disabled,
		Meta: &core.Meta{
			Title:              req.Meta.Title,
			Icon:               req.Meta.Icon,
			HideMenu:           req.Meta.HideMenu,
			HideBreadcrumb:     req.Meta.HideBreadcrumb,
			CurrentActiveMenu:  req.Meta.CurrentActiveMenu,
			IgnoreKeepAlive:    req.Meta.IgnoreKeepAlive,
			HideTab:            req.Meta.HideTab,
			FrameSrc:           req.Meta.FrameSrc,
			CarryParam:         req.Meta.CarryParam,
			HideChildrenInMenu: req.Meta.HideChildrenInMenu,
			Affix:              req.Meta.Affix,
			DynamicLevel:       req.Meta.DynamicLevel,
			RealPath:           req.Meta.RealPath,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}
