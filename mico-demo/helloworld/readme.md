# api demo
## 启动 Consul
```shell
# install
brew install consul
# run
consul agent -dev
```

## 编译proto
```shell
protoc --go_out=plugins=grpc:. --micro_out==plugins=grpc:. api.proto
protoc --go_out=plugins=grpc:. --micro_out==plugins=grpc:. api-base.proto
```

## 启动 api-getway
```bash
micro api --handler=api
```
## 启动web 管理页面
```bash
micro web
```

## 启动api服务
```bash
go run api.go
```

## 测试服务
```bash
curl -X GET \
  'http://localhost:8080/hello/example/call?name=johnasdfasdf' \
  -H 'Postman-Token: 8c60d5cd-639a-4cca-afae-2c67a8d83c6b' \
  -H 'cache-control: no-cache'
```

```bash
curl -X POST \
  http://localhost:8080/hello/foo/bar \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: eeca4f57-1242-4980-944c-c91f45610196' \
  -H 'cache-control: no-cache' \
  -d '{
	"name":"hahah"
}'
```
