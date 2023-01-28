package logic

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/message"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateTokenLogic {
	return &CreateOrUpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Token management
func (l *CreateOrUpdateTokenLogic) CreateOrUpdateToken(in *core.TokenInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		result := l.svcCtx.DB.Create(&model.Token{
			Model:    gorm.Model{},
			UUID:     in.UUID,
			Token:    in.Token,
			Status:   in.Status,
			Source:   in.Source,
			ExpireAt: time.Unix(in.ExpireAt, 0),
		})
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("Token already exists", logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, message.DictionaryAlreadyExists)
		}

		return &core.BaseResp{Msg: new_errorx.CreateSuccess}, nil
	} else {
		var origin model.Token
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if check.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}

		if check.RowsAffected == 0 {
			logx.Errorw(logmessage.TargetNotFound, logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, new_errorx.TargetNotExist)
		}

		origin.UUID = in.UUID
		origin.Token = in.Token
		origin.Status = in.Status
		origin.Source = in.Source
		origin.ExpireAt = time.Unix(in.ExpireAt, 0)

		result := l.svcCtx.DB.Save(&origin)

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			logx.Errorw(logmessage.UpdateFailed, logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
		}

		return &core.BaseResp{Msg: new_errorx.UpdateSuccess}, nil
	}
}
