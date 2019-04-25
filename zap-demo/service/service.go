package service

import (
	"time"

	"go.uber.org/zap"

	"github.com/akka/zap-demo/log"
)

func ShowSomething() {
	log.Info("service无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
