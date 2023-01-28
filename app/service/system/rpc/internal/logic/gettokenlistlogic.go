package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/message"
	"go-zero-pass/app/service/system/rpc/internal/model"
	"go-zero-pass/app/service/system/rpc/internal/svc"
	"go-zero-pass/app/service/system/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *core.TokenListReq) (*core.TokenListResp, error) {
	if in.Username == "" && in.Email == "" && in.Nickname == "" {
		var tokens []model.Token
		result := l.svcCtx.DB.Model(&model.Token{}).
			Limit(int(in.Page.PageSize)).Offset(int((in.Page.Page - 1) * in.Page.PageSize)).Find(&tokens)

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		resp := &core.TokenListResp{}
		resp.Total = uint64(result.RowsAffected)
		for _, v := range tokens {
			resp.Data = append(resp.Data, &core.TokenInfo{
				Id:       uint64(v.ID),
				UUID:     v.UUID,
				Token:    v.Token,
				Status:   v.Status,
				Source:   v.Source,
				ExpireAt: v.ExpireAt.UnixMilli(),
				CreateAt: v.CreatedAt.UnixMilli(),
			})
		}

		return resp, nil
	} else {
		var user model.User
		udb := l.svcCtx.DB.Model(&model.User{})

		if in.Username != "" {
			udb = udb.Where("username = ?", in.Username)
		}

		if in.Email != "" {
			udb = udb.Where("email = ?", in.Email)
		}

		if in.Nickname != "" {
			udb = udb.Where("nickname = ?", in.Nickname)
		}

		userData := udb.First(&user)

		if userData.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", userData.Error.Error()))
			return nil, status.Error(codes.Internal, userData.Error.Error())
		}

		if userData.RowsAffected == 0 {
			logx.Errorw(logmessage.TargetNotFound, logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, message.UserNotExists)
		}

		var tokens []model.Token
		result := l.svcCtx.DB.Where("UUID = ?", user.UUID).Limit(int(in.Page.PageSize)).
			Offset(int((in.Page.Page - 1) * in.Page.PageSize)).Find(&tokens)

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		resp := &core.TokenListResp{}
		resp.Total = uint64(result.RowsAffected)
		for _, v := range tokens {
			resp.Data = append(resp.Data, &core.TokenInfo{
				Id:       uint64(v.ID),
				UUID:     v.UUID,
				Token:    v.Token,
				Status:   v.Status,
				Source:   v.Source,
				ExpireAt: v.ExpireAt.UnixMilli(),
				CreateAt: v.CreatedAt.UnixMilli(),
			})
		}

		return resp, nil
	}
}
