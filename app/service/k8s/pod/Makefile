SHELL = /bin/bash


#生产api描述性文件
api-desc:
	goctl api -o ./api/desc/pod.api

api-code:
	goctl api go -api ./api/desc/pod.api -dir ./api/ -style gozero

#rpc
build: build-web build-server
	docker run --name build-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-local

