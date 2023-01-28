package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/response/new_errorx"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTokenStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetTokenStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTokenStatusLogic {
	return &SetTokenStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetTokenStatusLogic) SetTokenStatus(in *core.SetStatusReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Table("tokens").Where("id = ?", in.Id).Update("status", in.Status)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Update token status failed, please check the token id", logx.Field("TokenId", in.Id))
		return nil, status.Error(codes.InvalidArgument, new_errorx.UpdateFailed)
	}

	logx.Infow("Update token status successfully", logx.Field("TokenId", in.Id),
		logx.Field("Status", in.Status))
	return &core.BaseResp{Msg: new_errorx.UpdateSuccess}, nil
}
