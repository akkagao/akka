package main

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *lumberjack.Logger
)

func main() {
	logger = &lumberjack.Logger{
		Filename:   "./logs/demo.log", // 默认文件
		MaxSize:    1,                 // 最大文件尺寸（超过备份）
		MaxBackups: 60,                // 最大备份数量
		MaxAge:     30,                // 最大备份保存时长
		Compress:   true,              // 是否压缩
	}
	log.SetOutput(logger) // 接管系统日志输出
	log.Println("fsdf")   // 测试输出

}
