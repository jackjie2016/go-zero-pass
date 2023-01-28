package model

import "go-zero-pass/app/common/model"

type SvcPort struct {
	model.BaseModel
	SvcID int64 `json:"svc_id"`
	//服务的port
	SvcPort int32 `json:"svc_port"`
	//pod 中需要映射的port地址
	SvcTargetPort int32 `json:"svc_target_port"`
	//开启NodePort的模式下进行设置
	SvcNodePort int32 `json:"svc_node_port"`
	//端口协议
	SvcPortProtocol string `json:"svc_port_protocol"`
}
