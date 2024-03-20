package common

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
)

var Log = logrus.New()

type LumberHook struct {
	logger *lumberjack.Logger
}

func (hook LumberHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook LumberHook) Fire(entry *logrus.Entry) (err error) {

	fields := entry.Data
	var writeLine string
	if len(fields) != 0 {
		writeLine = fmt.Sprintf("[GIN] %s | %s | %d | %s | %s | %f \n",
			fields["time"],
			fields["method"],
			fields["status"],
			fields["path"],
			fields["clientIp"],
			fields["timeSub"],
		)
	} else {
		writeLine = entry.Message + "\n"
	}

	_, err = hook.logger.Write([]byte(writeLine))
	return
}

type MyFormatter struct{}

func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// 设置buffer 缓冲区
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	// 设置格式
	_, err := fmt.Fprintf(b, "%s\n", entry.Message)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func InitLog(logPath, appName string) {

	// 设置日志输出格式
	Log.SetFormatter(&MyFormatter{})

	logFileName := path.Join(logPath, appName)
	// 使用滚动压缩方式记录日志
	lumberLog := &lumberjack.Logger{
		Filename:   logFileName, //日志文件位置
		MaxSize:    1,           // 单文件最大容量,单位是MB
		MaxBackups: 3,           // 最大保留过期文件个数
		MaxAge:     1,           // 保留过期文件的最大时间间隔,单位是天
		Compress:   true,        // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}

	fileHook := LumberHook{lumberLog}

	Log.AddHook(&fileHook)
}
