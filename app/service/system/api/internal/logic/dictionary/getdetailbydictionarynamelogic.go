package dictionary

import (
	"context"
	"go-zero-pass/app/service/system/rpc/types/core"

	"go-zero-pass/app/service/system/api/internal/svc"
	"go-zero-pass/app/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailByDictionaryNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailByDictionaryNameLogic {
	return &GetDetailByDictionaryNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailByDictionaryNameLogic) GetDetailByDictionaryName(req *types.DictionaryDetailReq) (resp *types.DictionaryDetailListResp, err error) {
	result, err := l.svcCtx.CoreRpc.GetDetailByDictionaryName(context.Background(), &core.DictionaryDetailReq{Name: req.Name})

	if err != nil {
		return nil, err
	}

	resp = &types.DictionaryDetailListResp{}
	resp.Total = result.Total
	for _, v := range result.Data {
		resp.Data = append(resp.Data, types.DictionaryDetailInfo{
			Id:       v.Id,
			CreateAt: v.CreateAt,
			Title:    v.Title,
			Key:      v.Key,
			Value:    v.Value,
			Status:   v.Status,
		})
	}

	return resp, nil
}
