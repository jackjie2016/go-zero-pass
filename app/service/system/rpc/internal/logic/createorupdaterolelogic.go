package logic

import (
	"context"
	"fmt"
	"go-zero-pass/app/common/response/new_errorx"
	"strconv"
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

type CreateOrUpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateRoleLogic {
	return &CreateOrUpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// role service
func (l *CreateOrUpdateRoleLogic) CreateOrUpdateRole(in *core.RoleInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		data := &model.Role{
			Model:         gorm.Model{},
			Name:          in.Name,
			Value:         in.Value,
			DefaultRouter: in.DefaultRouter,
			Status:        in.Status,
			Remark:        in.Remark,
			OrderNo:       in.OrderNo,
		}
		result := l.svcCtx.DB.Create(&data)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("Role value had been used", logx.Field("Detail", data))
			return nil, status.Error(codes.InvalidArgument, message.DuplicateRoleValue)
		}

		err := l.UpdateRoleInfoInRedis()
		if err != nil {
			logx.Errorw("Fail to update the role info in Redis", logx.Field("Detail", err.Error()))
			return nil, err
		}

		logx.Infow("Create role successfully", logx.Field("Detail", data))
		return &core.BaseResp{Msg: new_errorx.CreateSuccess}, nil
	} else {
		var origin *model.Role
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if check.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}
		if check.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
		}
		data := &model.Role{
			Model:         gorm.Model{ID: origin.ID, CreatedAt: origin.CreatedAt, UpdatedAt: time.Now()},
			Name:          in.Name,
			Value:         in.Value,
			DefaultRouter: in.DefaultRouter,
			Status:        in.Status,
			Remark:        in.Remark,
			OrderNo:       in.OrderNo,
		}
		result := l.svcCtx.DB.Save(&data)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
		}
		err := l.UpdateRoleInfoInRedis()
		if err != nil {
			logx.Errorw("Fail to update the role info in Redis", logx.Field("Detail", err.Error()))
			return nil, err
		}

		logx.Infow("Update role successfully", logx.Field("Detail", data))
		return &core.BaseResp{Msg: new_errorx.UpdateSuccess}, nil
	}
}

func (l *CreateOrUpdateRoleLogic) UpdateRoleInfoInRedis() error {
	var roleData []model.Role
	res := l.svcCtx.DB.Find(&roleData)
	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, new_errorx.TargetNotExist)
	}
	for _, v := range roleData {
		err := l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", v.ID), v.Name)
		err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
		err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
		if err != nil {
			return status.Error(codes.Internal, new_errorx.RedisError)
		}
	}
	return nil
}
