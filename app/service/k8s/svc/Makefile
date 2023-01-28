SHELL = /bin/bash

#生产api描述性文件
api-desc:
	goctl api -o ./api/desc/svc.api

api-code:
	goctl api go -api ./api/desc/svc.api -dir ./api/ -style gozero

#rpc
rpc-proto:
	goctl rpc -o ./rpc/pb/svc.proto

rpc-code:
	goctl rpc protoc ./rpc/pb/svc.proto --go_out=. --go-grpc_out=. --zrpc_out=.

gen_model:
