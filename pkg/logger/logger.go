package logger

import (
	"github.com/sirupsen/logrus"
	"gostu/pkg/config"
	"gostu/pkg/logger/formatter"
	"gostu/pkg/logger/hook"
)

var Logger *logrus.Logger

func init()  {
	if Logger == nil {
		Logger = logrus.New()
		Logger.SetLevel(logrus.TraceLevel)

		// 添加响应事件
		setHook()
	}
}

func logFormatter(format string) logrus.Formatter {
	if format == "json" {
		return &logrus.JSONFormatter{TimestampFormat:   "2006-01-02 15:04:05"}
	} else if format == "customize_text" {
		return &formatter.CustLogFormatter{}
	} else {
		return &logrus.TextFormatter{TimestampFormat:   "2006-01-02 15:04:05"}
	}
}

func setHook()  {
	//日志输出格式
	format := config.AppConfig.GetString("log.logformat")
	// Info
	Logger.AddHook(hook.Info(logFormatter(format)))
	// Error
	Logger.AddHook(hook.Error(logFormatter(format)))
	// Debug
	if config.AppConfig.GetBool("debug") {
		Logger.AddHook(hook.Debug(logFormatter("customize_text")))
	}
}


