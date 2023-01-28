package svc

import (
	"github.com/casbin/casbin/v2"
	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/utils"
	"go-zero-pass/app/service/system/api/internal/config"
	"go-zero-pass/app/service/system/api/internal/middleware"
	"go-zero-pass/app/service/system/rpc/coreclient"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   coreclient.Core
	Redis     *redis.Redis
	Casbin    *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize redis
	rds := c.RedisConf.NewRedis()
	logx.Info("Initialize redis connection successfully")

	// initialize database connection
	db, err := c.DatabaseConf.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successful")

	// initialize casbin
	cbn := utils.NewCasbin(db)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		CoreRpc:   coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:     rds,
		Casbin:    cbn,
	}
}
