package model

import "go-zero-pass/app/common/model"

type PodPort struct {
	model.BaseModel
	PodID         int64  `json:"pod_id"`
	ContainerPort int32  `json:"container_port"`
	Protocol      string `json:"protocol"`
	//@TODO HostPort 需要的可以自主添加
}
