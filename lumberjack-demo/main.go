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
		Filename:   "./logs/demo.log",
		MaxSize:    1,
		MaxBackups: 60,
		MaxAge:     30,
		Compress:   true,
	}
	log.SetOutput(logger)
	log.Println("fsdf")

}
