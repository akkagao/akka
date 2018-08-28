安装
go get github.com/vaporz/turbo
编译
cd github.com/vaporz/turbo
make

创建项目
turbo create github.com/akkagao/turbo-demo/hello Hello -r grpc

进入hello后
go run main.go

调用接口
curl -X GET \
  'http://localhost:8081/hello?yourName=hahah' \
  -H 'Cache-Control: no-cache' \
  -H 'Postman-Token: 8a77ba8b-0153-4d4c-978e-1701f893c460'


重新生成存根
turbo generate github.com/akkagao/turbo-demo/hello -r grpc -I $GOPATH/src/github.com/akkagao/turbo-demo/hello


https://vaporz.github.io/master/zh/interceptor.html