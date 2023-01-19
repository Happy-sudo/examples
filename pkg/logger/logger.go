package logger

import (
	"examples/app/hello/internel/conf"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

//CuttingLogWriter 切割日志
func CuttingLogWriter(conf *conf.Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s", conf.Logger.Filename),
		MaxSize:    conf.Logger.MaxSize,    //日志的最大大小（M）
		MaxBackups: conf.Logger.MaxBackups, //日志的最大保存数量
		MaxAge:     conf.Logger.MaxAge,     //日志文件存储最大天数
		Compress:   conf.Logger.Compress,   //是否执行压缩
		LocalTime:  conf.Logger.LocalTime,  // 是否使用格式化时间辍
	}
	return zapcore.AddSync(lumberJackLogger)
}
