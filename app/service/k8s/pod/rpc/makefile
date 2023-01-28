
host  = localhost
port  = 3306
username  = root
passwd  = root123
tables  = root
modeldir =../model
dbname = simple_admin

#容器环境打包前端
Pod_model:
	goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="pod"  -dir="${modeldir}" -cache=false --style=goZero
