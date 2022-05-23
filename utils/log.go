package utils

import (
	"io"
	"law/conf"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var appBufFile *BufWriter

func init() {
	if conf.TestMode {
		setNullLogger()
		return
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if conf.App.IsDev() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	appFile, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件app.log")
	}
	appBufFile = NewBufWriter(appFile)
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
