// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Id int64 `path:"Id"`
}

type Response struct {
}

type BaseMsg struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type BaseResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type FindAll struct {
}

type PodList struct {
	PodInfo []PodInfo `json:"pod_info"`
}

type PodInfo struct {
	Id            int64     `json:"id"`
	PodNamespace  string    `json:"pod_namespace"`
	PodName       string    `json:"pod_name"`
	PodTeamId     string    `json:"pod_team_id"`
	PodCpuMax     float32   `json:"pod_cpu_max"`
	PodReplicas   int32     `json:"pod_replicas"`
	PodMemoryMax  float32   `json:"pod_memory_max"`
	PodPort       []PodPort `json:"pod_port"`
	PodEnv        []PodEnv  `json:"pod_env"`
	PodPullPolicy string    `json:"pod_pull_policy"`
	PodRestart    string    `json:"pod_restart"`
	PodType       string    `json:"pod_type"`
	PodImage      string    `json:"pod_image"`
}

type PodPort struct {
	PodId         int64  `json:"pod_id"`
	ContainerPort int32  `json:"container_port"`
	Protocol      string `json:"protocol"`
}

type PodEnv struct {
	PodId    int64  `json:"pod_id"`
	EnvKey   string `json:"env_key"`
	EnvValue string `json:"env_value"`
}