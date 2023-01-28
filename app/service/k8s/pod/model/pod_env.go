package model

import "go-zero-pass/app/common/model"

type PodEnv struct {
	model.BaseModel
	PodID    int64  `json:"pod_id"`
	EnvKey   string `json:"env_key"`
	EnvValue string `json:"env_value"`
}
