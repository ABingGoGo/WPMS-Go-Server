package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

const FilesNextName = ".log"

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		urlPath := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		// Stop timer
		end := time.Now()
		timeSub := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		if raw != "" {
			urlPath = urlPath + "?" + raw
		}

		logrus.Infof("[GIN]-time:%s StatusCode:%d ip:%s 耗时:%d 访问路径:%s 访问动作:%s",
			start.Format("2006-01-02 15:04:06"),
			statusCode,
			clientIP,
			timeSub,
			urlPath,
			method,
		)

	}
}

type MyFormatter struct {
}

func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var msg string
	if entry.HasCaller() {
		msg = fmt.Sprintf("%s %s [%s] [%d] - %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Caller.Function, entry.Caller.Line, entry.Message)
	} else {
		msg = fmt.Sprintf("%s [%s] - %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	}
	return []byte(msg), nil
}

func writer(logPath string, level string /*, save uint*/) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logPath, level)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fileSuffix := time.Now().In(cstSh).Format("2006-01-02") + FilesNextName

	logier, err := rotatelogs.New(
		logFullPath+"-"+fileSuffix,
		rotatelogs.WithLinkName(logFullPath), // 生成软链，指向最新日志文件
		// rotatelogs.WithRotationCount(save),        // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 设置最大保存时间(7天)
	)

	if err != nil {
		panic(err)
	}
	return logier
}

func InitFile(logPath string) {
	logrus.SetFormatter(&MyFormatter{})

	fileDate := time.Now().Format("2006-01-02")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}
	fileHoo := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer(fmt.Sprintf("%s/%s", logPath, fileDate), "debug" /*, 8*/), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer(fmt.Sprintf("%s/%s", logPath, fileDate), "info" /*, 8*/),
		logrus.WarnLevel:  writer(fmt.Sprintf("%s/%s", logPath, fileDate), "warn" /*, 8*/),
		logrus.ErrorLevel: writer(fmt.Sprintf("%s/%s", logPath, fileDate), "error" /*, 8*/),
		logrus.FatalLevel: writer(fmt.Sprintf("%s/%s", logPath, fileDate), "fatal" /*, 8*/),
		logrus.PanicLevel: writer(fmt.Sprintf("%s/%s", logPath, fileDate), "panic" /*, 8*/),
	}, &MyFormatter{})

	logrus.AddHook(fileHoo)
}
