package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-pass/app/common/logmessage"
	"go-zero-pass/app/common/utils/k8s"
	"go-zero-pass/app/service/k8s/volume/model/repository"
	"go-zero-pass/app/service/k8s/volume/model/service"
	"go-zero-pass/app/service/k8s/volume/rpc/internal/config"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
)

type ServiceContext struct {
	Config       config.Config
	DbClient     *gorm.DB
	K8sClient    *kubernetes.Clientset
	ModelService service.IVolumeDataService
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := c.DB.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successfully")
	//db.AutoMigrate(&model.Svc{}, &model.SvcPort{})
	volumeRepository := repository.NewVolumeRepository(db)
	if err := volumeRepository.InitTable(); err != nil {
		logx.Errorw(logmessage.DBInitError, logx.Field("Detail", err.Error()))
		return nil
	}

	K8sClient, err := k8s.NewK8sClient()
	if err != nil {
		logx.Errorw(logmessage.K8SError, logx.Field("Detail", err.Error()))
		return nil
	}
	logx.Info("Initialize k8s connection successfully")
	return &ServiceContext{
		Config:       c,
		DbClient:     db,
		K8sClient:    K8sClient,
		ModelService: service.NewVolumeDataService(volumeRepository, K8sClient),
	}
}
