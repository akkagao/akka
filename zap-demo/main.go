package main

import (
	"time"

	"go.uber.org/zap"

	"github.com/akka/zap-demo/log"
	"github.com/akka/zap-demo/service"
)

func main() {

	if log.DebugEnabled() {
		log.Info("无法获取网址",
			zap.String("url", "http://www.baidu.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
	service.ShowSomething()

}
