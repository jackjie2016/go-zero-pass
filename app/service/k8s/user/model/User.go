package model

import "go-zero-pass/app/common/model"

type User struct {
	model.BaseModel
	UserName   string  `gorm:"not_null;unique" json:"user_name"`
	UserEmail  string  `gorm:"not_null;unique" json:"user_email"`
	IsAdmin    *bool   `json:"is_admin;omitempty"`
	UserPwd    string  `json:"user_pwd"`
	UserStatus int32   `json:"user_status"`
	Role       []*Role `gorm:"many2many:user_role"`
}
