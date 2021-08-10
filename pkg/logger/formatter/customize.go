package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

//自定义日志格式
type CustLogFormatter struct {}

func (f *CustLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	content, _ := json.Marshal(entry.Data)
	msg := fmt.Sprintf("[%s] %s %v\n", timestamp, entry.Message, string(content))
	return []byte(msg), nil
}
