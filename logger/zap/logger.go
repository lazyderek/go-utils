package zap

import (
	"github.com/lazyderek/config/yaml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const (
	LevelInfo  = "info"
	LevelDebug = "debug"
	LevelError = "error"
)

var log Logger

type Logger = *logger

type logger struct {
	*zap.Logger
}

func init() {
	Init(LevelInfo, "")
}

func Init(logLevel string, logPath string) Logger {
	var (
		encoderConfig = zap.NewProductionEncoderConfig()
		core          zapcore.Core
		level         = getLogLevel(logLevel)
	)

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if logPath == "" {
		consoleDebugging := zapcore.Lock(os.Stdout)
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			consoleDebugging, // 输出到控制台
			level,
		)
	} else {
		hook := getWriter(logPath)
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(hook), // 输入文本
			level,
		)
	}

	l := zap.New(core)

	log = &logger{Logger: l}
	log.Info("logger init...")
	return log
}

func Info(msg string, fields ...zap.Field) {
	log.Logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Logger.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Logger.Warn(msg, fields...)
}

func getLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case LevelDebug:
		return zap.DebugLevel
	case LevelInfo:
		return zap.InfoLevel
	case LevelError:
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getWriter(filename string) io.Writer {
	hook, _ := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 切割频率 24小时
	)
	return hook
}
