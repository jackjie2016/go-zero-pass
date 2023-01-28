package logic

import (
	"context"
	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/message"
	"go-zero-pass/app/common/response/new_errorx"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryLogic {
	return &CreateOrUpdateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// dictionary management service
func (l *CreateOrUpdateDictionaryLogic) CreateOrUpdateDictionary(in *core.DictionaryInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		result := l.svcCtx.DB.Create(&model.Dictionary{
			Model:  gorm.Model{},
			Title:  in.Title,
			Name:   in.Name,
			Status: in.Status,
			Desc:   in.Desc,
			Detail: nil,
		})
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("Dictionary already exists", logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, message.DictionaryAlreadyExists)
		}

		return &core.BaseResp{Msg: new_errorx.CreateSuccess}, nil
	} else {
		var origin model.Dictionary
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if check.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}

		if check.RowsAffected == 0 {
			logx.Errorw(logmessage.TargetNotFound, logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, new_errorx.TargetNotExist)
		}

		origin.Title = in.Title
		origin.Name = in.Name
		origin.Status = in.Status
		origin.Desc = in.Desc

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
