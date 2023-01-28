package logic

import (
	"context"
	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/response/new_errorx"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryDetailLogic {
	return &DeleteDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryDetailLogic) DeleteDictionaryDetail(in *core.IDReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Delete(&model.DictionaryDetail{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Delete dictionary detail failed, check the id", logx.Field("DetailId", in.ID))
		return nil, status.Error(codes.InvalidArgument, new_errorx.DeleteFailed)
	}

	logx.Infow("Delete dictionary detail successfully", logx.Field("DetailId", in.ID))

	return &core.BaseResp{Msg: new_errorx.DeleteSuccess}, nil
}
