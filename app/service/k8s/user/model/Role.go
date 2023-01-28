package model

import "go-zero-pass/app/common/model"

type Role struct {
	model.BaseModel
	RoleName   string        `json:"role_name"`
	RoleStatus int32         `json:"role_status"`
	Permission []*Permission `gorm:"many2many:role_permission`
}
