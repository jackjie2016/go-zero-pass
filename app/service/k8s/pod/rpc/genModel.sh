#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$1
#表生成的genmodel目录
modeldir=../model

# 数据库配置
host=127.0.0.1
port=3306
dbname=simple_admin
username=root
passwd=root123


echo "开始创建库：$dbname 的表：$1"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=false --style=goZero