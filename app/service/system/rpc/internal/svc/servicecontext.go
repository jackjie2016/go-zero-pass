package svc

import (
	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/service/system/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize database connection
	db, err := c.DB.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successfully")
	// initialize redis
	rds := c.RedisConf.NewRedis()
	logx.Info("Initialize redis connection successfully")
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
	}
}
