package logic

import (
	"context"
	"fmt"
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

type DeleteDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(in *core.IDReq) (*core.BaseResp, error) {
	childResult := l.svcCtx.DB.Exec(fmt.Sprintf("delete from dictionary_details where dictionary_id = %d", in.ID))
	if childResult.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", childResult.Error.Error()))
		return nil, status.Error(codes.Internal, childResult.Error.Error())
	}

	result := l.svcCtx.DB.Delete(&model.Dictionary{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Delete dictionary failed, check the id", logx.Field("DictionaryId", in.ID))
		return nil, status.Error(codes.InvalidArgument, new_errorx.DeleteFailed)
	}

	logx.Infow("Delete dictionary successfully", logx.Field("DictionaryId", in.ID))

	return &core.BaseResp{Msg: new_errorx.DeleteSuccess}, nil
}
