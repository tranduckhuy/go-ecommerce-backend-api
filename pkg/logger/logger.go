package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/tranduckhuy/go-ecommerce-backend-api/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config settings.LogConfig) *LoggerZap {
	var level zapcore.Level
	err := level.UnmarshalText([]byte(config.Level))
	if err != nil {
		level = zapcore.InfoLevel
	}

	encoder := getEncoder(config.Encoding)
	writer := getLogWriter(config)

	core := zapcore.NewCore(encoder, writer, level)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoder(encoding string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	if encoding == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(config settings.LogConfig) zapcore.WriteSyncer {
	hook := lumberjack.Logger{
		Filename:   config.OutputPath,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
		Compress:   config.Compress,
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
}
