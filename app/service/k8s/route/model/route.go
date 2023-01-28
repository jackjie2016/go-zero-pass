package model

import "go-zero-pass/app/common/model"

type Route struct {
	model.BaseModel
	RouteName      string      `json:"route_name"`
	RouteNamespace string      `json:"route_namespace"`
	RouteHost      string      `json:"route_host"`
	RoutePath      []RoutePath `gorm:"ForeignKey:RouteID" json:"route_path"`
}
