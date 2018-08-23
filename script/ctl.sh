#!/bin/sh

start(){
    cd ../server
    go run main.go
}

proto(){
    cd ../proto/
    ./export.sh
}

config(){
    cd ../../acinconfig/bin
    ./acinconfig
}

help()
{
	echo " manager command:     "
	echo " start    以正常服务器方式启动  "
	echo " stop     先踢人下线再关闭服务器  "
	echo " remote   远程连接服务器  "
	echo " proto    生成protobuff文件  "
	echo " config   生成配置文件  "
}

case $1 in
	'start') start ;;
	'stop') stop ;;
	'remote') remote ;;
	'proto') proto ;;
	'config') config ;;
	*) help ;;
esac