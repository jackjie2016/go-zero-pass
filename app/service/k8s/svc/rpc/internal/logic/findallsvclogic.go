package logic

import (
	"context"
	common "go-zero-pass/app/common/utils"

	"go-zero-pass/app/service/k8s/svc/rpc/internal/svc"
	"go-zero-pass/app/service/k8s/svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllSvcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllSvcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllSvcLogic {
	return &FindAllSvcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllSvcLogic) FindAllSvc(in *pb.FindAll) (*pb.AllSvc, error) {
	// todo: add your logic here and delete this line
	logx.Info("查询所有服务")
	Svcs, err := l.svcCtx.ModelService.FindAllSvc()
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	AllSvc := pb.AllSvc{}
	//整理格式
	for _, v := range Svcs {
		svcInfo := &pb.SvcInfo{}
		if err := common.SwapTo(v, svcInfo); err != nil {
			logx.Error(err)
			return nil, err
		}
		AllSvc.SvcInfo = append(AllSvc.SvcInfo, svcInfo)
	}

	return &AllSvc, nil
}
