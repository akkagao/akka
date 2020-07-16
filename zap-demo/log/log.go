package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global variables against which all our logging occurs.
var logger *zap.Logger
//var sugar = logger.Sugar()

func init() {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短径编码器
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.InfoLevel)

	config := zap.Config{
		Level:            atom,                                              // 日志级别
		Development:      true,                                              // 开发模式，堆栈跟踪
		Encoding:         "console",                                         // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                     // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "zap-demo"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout", "./logs/spikeProxy.log"},       // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	l, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
	logger = l

}

// Debug outputs a message at debug level.
// This call is a wrapper around [Logger.Debug](https://godoc.org/go.uber.org/zap#Logger.Debug)
func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

// DebugEnabled returns whether output of messages at the debug level is currently enabled.
func DebugEnabled() bool {
	return logger.Core().Enabled(zap.DebugLevel)
}

// Error outputs a message at error level.
// This call is a wrapper around [logger.Error](https://godoc.org/go.uber.org/zap#logger.Error)
func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

// ErrorEnabled returns whether output of messages at the error level is currently enabled.
func ErrorEnabled() bool {
	return logger.Core().Enabled(zap.ErrorLevel)
}

// Warn outputs a message at warn level.
// This call is a wrapper around [logger.Warn](https://godoc.org/go.uber.org/zap#logger.Warn)
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

// WarnEnabled returns whether output of messages at the warn level is currently enabled.
func WarnEnabled() bool {
	return logger.Core().Enabled(zap.WarnLevel)
}

// Info outputs a message at information level.
// This call is a wrapper around [logger.Info](https://godoc.org/go.uber.org/zap#logger.Info)
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

// InfoEnabled returns whether output of messages at the info level is currently enabled.
func InfoEnabled() bool {
	return logger.Core().Enabled(zap.InfoLevel)
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
// This call is a wrapper around [logger.With](https://godoc.org/go.uber.org/zap#logger.With)
func With(fields ...zapcore.Field) *zap.Logger {
	return logger.With(fields...)
}

// Sync flushes any buffered log entries.
// Processes should normally take care to call Sync before exiting.
// This call is a wrapper around [logger.Sync](https://godoc.org/go.uber.org/zap#logger.Sync)
func Sync() error {
	return logger.Sync()
}
