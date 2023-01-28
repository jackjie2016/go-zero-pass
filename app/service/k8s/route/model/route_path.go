package model

import "go-zero-pass/app/common/model"

type RoutePath struct {
	model.BaseModel
	RouteID                 int64  `json:"route_id"`
	RoutePathName           string `json:"route_path_name"`
	RouteBackendService     string `json:"route_backend_service"`
	RouteBackendServicePort int32  `json:"route_backend_service_port"`
}
