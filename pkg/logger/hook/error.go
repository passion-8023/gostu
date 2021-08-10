package hook

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gostu/pkg/config"
	"time"
)

func Error(formatter logrus.Formatter) *lfshook.LfsHook {
	writer, _ := rotatelogs.New(
		config.AppConfig.GetString("log.logpath") + "error/error_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Hour * 24 * config.AppConfig.GetDuration("log.max_age")),
		rotatelogs.WithRotationTime(time.Hour),
	)
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,

	}, formatter)
}
