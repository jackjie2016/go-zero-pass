SHELL = /bin/bash

#生产api描述性文件
api-desc:
	goctl api -o ./api/desc/user.api

api-code:
	goctl api go -api ./api/desc/user.api -dir ./api/ -style gozero

#rpc
rpc-proto:
	goctl rpc -o ./rpc/pb/user.proto

rpc-code:
	goctl rpc protoc ./rpc/pb/user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

gen_model:
