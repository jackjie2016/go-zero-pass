package logic

import (
	"context"
	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/response/new_errorx"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeMenuParamListByMenuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeMenuParamListByMenuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeMenuParamListByMenuIdLogic {
	return &GeMenuParamListByMenuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeMenuParamListByMenuIdLogic) GeMenuParamListByMenuId(in *core.IDReq) (*core.MenuParamListResp, error) {
	var paramsList []model.MenuParam
	result := l.svcCtx.DB.Where("menu_id = ?", in.ID).Find(&paramsList)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, new_errorx.DatabaseError)
	}

	resp := &core.MenuParamListResp{}
	resp.Total = uint64(result.RowsAffected)
	for _, v := range paramsList {
		resp.Data = append(resp.Data, &core.MenuParamResp{
			Id:       uint64(v.ID),
			MenuId:   uint64(v.MenuId),
			Type:     v.Type,
			Key:      v.Key,
			Value:    v.Value,
			CreateAt: v.CreatedAt.UnixMilli(),
			UpdateAt: v.UpdatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
