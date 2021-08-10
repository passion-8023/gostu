package hook

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gostu/pkg/config"
	"time"
)

func Info(formatter logrus.Formatter) *lfshook.LfsHook {
	writer, _ := rotatelogs.New(
		config.AppConfig.GetString("log.logpath")+"info/run_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Hour*24*config.AppConfig.GetDuration("log.max_age")),
		rotatelogs.WithRotationTime(time.Hour),
	)
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel: writer,
		logrus.WarnLevel: writer,
	}, formatter)
}

