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

type GetMenuListByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole(in *core.IDReq) (*core.MenuInfoList, error) {
	var r model.Role
	result := l.svcCtx.DB.Preload("Menus").Preload("Menus.Children", func(db *gorm.DB) *gorm.DB {
		return db.Order("menus.order_no DESC")
	}).Where(&model.Role{Model: gorm.Model{ID: uint(in.ID)}}).First(&r)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, "database error")
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "data not found")
	}
	var res *core.MenuInfoList
	res = &core.MenuInfoList{}
	res.Total = uint64(result.RowsAffected)
	var validId map[uint]struct{}
	validId = make(map[uint]struct{})
	for _, v := range r.Menus {
		validId[v.ID] = struct{}{}
	}
	res.Data = findRoleMenuChildren(r.Menus, validId, 1)
	// delete menus whose menu levels are not 1
	//var tmp []*core.MenuInfo
	//for _, v := range res.Data {
	//	if v.Level == 1 {
	//		tmp = append(tmp, v)
	//	}
	//}
	//res.Data = tmp
	return res, nil
}

func findRoleMenuChildren(data []model.Menu, validId map[uint]struct{}, parentId uint) []*core.MenuInfo {
	if data == nil {
		return nil
	}
	var result []*core.MenuInfo
	for _, v := range data {
		if v.ParentId == parentId && v.ID != v.ParentId {
			if _, ok := validId[v.ID]; ok {
				tmp := &core.MenuInfo{
					Id:        uint64(v.ID),
					CreateAt:  v.CreatedAt.UnixMilli(),
					UpdateAt:  v.UpdatedAt.UnixMilli(),
					MenuType:  v.MenuType,
					Level:     v.MenuLevel,
					ParentId:  uint32(v.ParentId),
					Path:      v.Path,
					Name:      v.Name,
					Redirect:  v.Redirect,
					Component: v.Component,
					OrderNo:   v.OrderNo,
					Meta: &core.Meta{
						Title:              v.Meta.Title,
						Icon:               v.Meta.Icon,
						HideMenu:           v.Meta.HideMenu,
						HideBreadcrumb:     v.Meta.HideBreadcrumb,
						CurrentActiveMenu:  v.Meta.CurrentActiveMenu,
						IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
						HideTab:            v.Meta.HideTab,
						FrameSrc:           v.Meta.FrameSrc,
						CarryParam:         v.Meta.CarryParam,
						HideChildrenInMenu: v.Meta.HideChildrenInMenu,
						Affix:              v.Meta.Affix,
						DynamicLevel:       v.Meta.DynamicLevel,
						RealPath:           v.Meta.RealPath,
					},
					Children: findRoleMenuChildren(data, validId, v.ID),
				}
				result = append(result, tmp)
			}
		}
	}
	return result
}
