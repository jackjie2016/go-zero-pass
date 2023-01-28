package role

import (
	"context"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.PageInfo) (resp *types.RoleListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetRoleList(context.Background(), &core.PageInfoReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.RoleListResp{}
	resp.Total = data.Total
	for _, v := range data.Data {
		resp.Data = append(resp.Data, types.RoleInfo{
			Id:            v.Id,
			Name:          v.Name,
			Value:         v.Value,
			DefaultRouter: v.DefaultRouter,
			Status:        v.Status,
			Remark:        v.Remark,
			OrderNo:       v.OrderNo,
			CreateAt:      v.CreateAt,
		})
	}
	return resp, nil
}
