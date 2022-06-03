package utils

import (
	"bufio"
	"io"
	"law/conf"
	"os"
	"time"

	stdLog "log"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
)

var appBufFile *bufio.Writer

func init() {
	//因为log会被很早加载，所以将时区设置放在这里
	initTimezone()
	if conf.TestMode {
		setNullLogger()
		return
	}
	//zerolog是本项目的主要日志库，
	//其他日志组件用于记录部分依赖库的日志
	initZerolog()
	///初始化标准日志组件
	initStdLog()
	//初始化logrus日志组件
	initLogrus()
}

func initZerolog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if conf.App.IsDev() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	appFile, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件app.log")
	}
	appBufFile = bufio.NewWriter(appFile)
	errFile, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件error.log")
	}
	lw := &LevelWriter{Writer: appBufFile, ErrorWriter: errFile}
	log.Logger = log.Output(lw)
	go func() {
		for {
			time.Sleep(time.Second * 5)
			FlushLog()
		}
	}()

}

func setNullLogger() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	logWriter, err := os.OpenFile(os.DevNull, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件" + os.DevNull)
	}
	log.Logger = log.Output(logWriter)
}

//由于bufio会缓存日志,可在需要时主动将缓存刷出
func FlushLog() {
	if appBufFile.Buffered() != 0 {
		appBufFile.Flush()
	}
}

type LevelWriter struct {
	io.Writer
	ErrorWriter io.Writer
}

func (lw *LevelWriter) WriteLevel(l zerolog.Level, p []byte) (n int, err error) {
	w := lw.Writer
	if zerolog.NoLevel > l && l > zerolog.InfoLevel {
		FlushLog() //将缓存的低级别日志刷盘
		w = lw.ErrorWriter
	}
	return w.Write(p)
}

func initTimezone() {
	loc, _ := time.LoadLocation("Asia/Shanghai") //加载时区
	time.Local = loc
}

func initStdLog() {
	standardLogFilePath := "logs/standard.log"
	logFile, err := os.OpenFile(standardLogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件" + standardLogFilePath)
	}
	stdLog.SetOutput(logFile)
}

func initLogrus() {
	logrusFilePath := "logs/logrus.log"
	logFile, err := os.OpenFile(logrusFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件" + logrusFilePath)
	}
	logrus.SetOutput(logFile)
}
