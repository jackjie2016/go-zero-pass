package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-pass/app/common/utils/gormsql"
)

type Config struct {
	zrpc.RpcServerConf
	DB        gormsql.GORMConf `json:"DatabaseConf" yaml:"DatabaseConf"`
	RedisConf redis.RedisConf  `json:"RedisConf" yaml:"RedisConf"`
	//K8sClinet kubernetes.Clientset `json:"K8s_clinet" yaml:"K8s_clinet"`
}
