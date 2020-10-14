package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
	"strings"
	"test/models"
	"time"
)


func LogMiddle(ctx iris.Context) {
	appName := viper.GetString("app_name")
	traceLogger := models.NewLogger()
	traceLogger.Proto = ctx.Request().Proto
	traceLogger.Referer = ctx.Request().Referer()
	traceLogger.ReqTime = time.Now()
	traceLogger.Length = ctx.Request().ContentLength
	traceLogger.UserAgent = ctx.Request().UserAgent()
	traceLogger.ReqUri = ctx.Request().RequestURI
	traceLogger.ReqMethod = ctx.Request().Method
	logrus.AddHook(NewContextHook(appName, traceLogger))
	logrus.Info("这是一条测试")
	ctx.Next()
}

/**
添加一个自定义的钩子，实现在日志中增加一个字段appName和请求的基本信息，
及其日志rotate策略
*/

type AppHook struct {
	Skip        int
	AppName     string
	TraceLogger *models.TraceLogger
}

func NewContextHook(appName string, logger *models.TraceLogger) *AppHook {
	return &AppHook{AppName: appName, TraceLogger: logger, Skip: 5}
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	entry.Data["proto"] = h.TraceLogger.Proto
	entry.Data["method"] = h.TraceLogger.ReqMethod
	entry.Data["uri"] = h.TraceLogger.ReqUri
	entry.Data["user-agent"] = h.TraceLogger.UserAgent
	entry.Data["length"] = h.TraceLogger.Length
	entry.Data["referer"] = h.TraceLogger.Referer
	entry.Data["line"] = findCaller(h.Skip)
	return nil
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}

func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}