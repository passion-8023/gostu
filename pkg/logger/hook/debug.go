package hook

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gostu/pkg/config"
	"time"
)

func Debug(formatter logrus.Formatter) *lfshook.LfsHook {
	writer, _ := rotatelogs.New(
		config.AppConfig.GetString("log.logpath") + "debug_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Hour * 2),
		rotatelogs.WithRotationTime(time.Hour),
	)

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.TraceLevel: writer,
	}, formatter)
}
