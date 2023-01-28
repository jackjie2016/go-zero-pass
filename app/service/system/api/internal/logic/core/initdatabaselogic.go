package core

import (
	"context"
	"go-zero-pass/app/common/response/new_errorx"
	"go-zero-pass/app/service/system/rpc/types/core"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.InitDatabase(l.ctx, &core.Empty{})
	if err != nil && !errors.Is(err, status.Error(codes.DeadlineExceeded, "context deadline exceeded")) {
		return nil, err
	} else if errors.Is(err, status.Error(codes.DeadlineExceeded, "context deadline exceeded")) {
		for {
			// wait 10 second for initialization
			time.Sleep(time.Second * 10)
			if initState, err := l.svcCtx.Redis.Get("database_init_state"); err == nil {
				if initState == "1" {
					return &types.SimpleMsg{Msg: new_errorx.Success}, nil
				}
			} else {
				return nil, status.Error(codes.Internal, new_errorx.RedisError)
			}

			if errMsg, err := l.svcCtx.Redis.Get("database_error_msg"); err == nil {
				if errMsg != "" {
					return nil, status.Error(codes.Internal, errMsg)
				}
			} else {
				return nil, status.Error(codes.Internal, new_errorx.RedisError)
			}
		}
	}
	return &types.SimpleMsg{Msg: result.Msg}, nil
}
