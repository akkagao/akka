#!/bin/bash

echo "清理编译结果"
rm main docker-go
echo "编译项目"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
echo "修改二进制文件名称"
mv main docker-go

echo "如果docker 已经启动则关闭docker"
dockerId=`docker ps -a | grep docker-go:latest | awk '{print $1}'`
if [ -n "$dockerId" ]; then
    docker stop $dockerId
    docker rm $dockerId
fi

echo "重新编译docker"
docker build -t docker-go:latest .
echo "启动docker服务"
docker run -d -p 8088:8088 docker-go:latest