使用docker启动rabbitmq

docker run -d --hostname my-rabbit --name some-rabbit rabbitmq:3
docker run -d  -p 5671:5671 -p 5672:5672  -p 15672:15672 -p 15671:15671  -p 25672:25672 rabbitmq