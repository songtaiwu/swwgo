package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"swwgo/mframe/package/setting/conf"
)

var errorLogger *zap.SugaredLogger


func defaultStr(v , fallback string) string {
	if v != "" {
		return v
	}
	return fallback
}

func defaultInt(v, fallback int) int {
	if v != 0 {
		return v
	}
	return fallback
}


var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	lvl = strings.ToLower(lvl)
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Init()  {
	// 从conf中读取配置
	filePath := defaultStr(conf.LogSetting.FilePath, "./logs")
	maxSize := defaultInt(conf.LogSetting.MaxSize, 128)
	maxBackups := defaultInt(conf.LogSetting.MaxBackups, 100)
	maxAge := defaultInt(conf.LogSetting.MaxAge, 30)
	level := defaultStr(conf.LogSetting.Level, "info")

	// 文件路径处理，
	// 如果是"/"结果，当做目录
	fileName := "log.log"
	if filePath != "" {
		if strings.HasSuffix(filePath, "/") {
			filePath = filePath + fileName
		} else {
			filePath = filePath + "/" + fileName
		}
	}

	// 日志等级
	logLevel := getLoggerLevel(level)

	// 日志writer
	l := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：Mb
		MaxAge:     maxAge,     // 文件最多保存多少天
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		LocalTime:  false,
		Compress:   true, // 是否压缩
	}
	syncWriter := zapcore.AddSync(&l)

	// 日志级别大于等于error时返回true
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})
	// 日志级别小于error时返回true
	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.ErrorLevel
	})

	// 生产日志格式
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	// 开发模式的输出格式
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(logLevel)),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}