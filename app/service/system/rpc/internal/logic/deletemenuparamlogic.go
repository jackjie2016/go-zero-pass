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

type DeleteMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuParamLogic {
	return &DeleteMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuParamLogic) DeleteMenuParam(in *core.IDReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Delete(&model.MenuParam{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Delete Menu parameter failed, check the id", logx.Field("MenuParamId", in.ID))
		return nil, status.Error(codes.InvalidArgument, new_errorx.DeleteFailed)
	}

	logx.Infow("Delete Menu parameter successfully", logx.Field("MenuParamId", in.ID))
	return &core.BaseResp{Msg: new_errorx.DeleteSuccess}, nil
}
