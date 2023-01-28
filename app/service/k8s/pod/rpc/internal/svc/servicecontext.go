package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/service/k8s/pod/model/repository"
	"go-zero-pass/app/service/k8s/pod/model/service"
	"go-zero-pass/app/service/k8s/pod/rpc/internal/config"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"

	"go-zero-pass/app/common/utils/k8s"
)

type ServiceContext struct {
	Config       config.Config
	DbClient     *gorm.DB
	K8sClient    *kubernetes.Clientset
	ModelService service.IPodDataService
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize database connection
	db, err := c.DB.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		return nil
	}
	PodRepository := repository.NewPodRepository(db)
	//db.AutoMigrate(&model.Pod{}, &model.PodPort{}, &model.PodEnv{})
	logx.Info("Initialize database connection successfully")

	K8sClient, err := k8s.NewK8sClient()
	if err != nil {
		logx.Errorw(logmessage.K8SError, logx.Field("Detail", err.Error()))
		return nil
	}
	logx.Info("Initialize k8s connection successfully")

	//// initialize redis
	//rds := c.RedisConf.NewRedis()
	//logx.Info("Initialize redis connection successfully")
	return &ServiceContext{
		Config:       c,
		DbClient:     db,
		K8sClient:    K8sClient,
		ModelService: service.NewPodDataService(PodRepository, K8sClient),
	}
}
