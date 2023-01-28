package model

import "go-zero-pass/app/common/model"

type Permission struct {
	model.BaseModel
	PermissionName     string `json:"permission_name"`
	PermissionDescribe string `json:"permission_describe"`
	PermissionAction   string `json:"permission_action"`
	PermissionStatus   int32  `json:"permission_status"`
}
