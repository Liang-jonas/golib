package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

type FormatterImpl interface {
	Format(entry *logrus.Entry) ([]byte, error)
}

type logFormatter struct {
	enableLogColor bool
}

func mapLevel(m logrus.Level) string {
	switch m {
	case logrus.TraceLevel:
		return "TRACE"
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.WarnLevel:
		return "WARNING"
	case logrus.ErrorLevel:
		return "ERROR"
	case logrus.FatalLevel:
		return "FATAL"
	case logrus.PanicLevel:
		return "PANIC"
	}
	return "UNKNOWN"
}

func mapLevelWithColor(m logrus.Level) string {
	switch m {
	case logrus.TraceLevel:
		return "\x1b[94mTRACE\x1b[0m"
	case logrus.DebugLevel:
		return "\x1b[36mDEBUG\x1b[0m"
	case logrus.InfoLevel:
		return "\x1b[32mINFO\x1b[0m"
	case logrus.WarnLevel:
		return "\x1b[33mWARNING\x1b[0m"
	case logrus.ErrorLevel:
		return "\x1b[31mERROR\x1b[0m"
	case logrus.FatalLevel:
		return "\x1b[35mFATAL\x1b[0m"
	case logrus.PanicLevel:
		return "\x1b[41mPANIC\x1b[0m"
	}
	return "UNKNOWN"
}
func getGID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}

func (l *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	if l.enableLogColor {
		return []byte(fmt.Sprintf("%s \x1b[90m%s:%d\x1b[0m %s [%s] %s\n", timestamp, file, len, getGID(), mapLevelWithColor(entry.Level), entry.Message)), nil
	}
	return []byte(fmt.Sprintf("%s %s:%d %s [%s] %s\n", timestamp, file, len, getGID(), mapLevel(entry.Level), entry.Message)), nil
}

func NewLogger(enableLogColor bool, logPath string) (*logrus.Logger, error) {
	now := time.Now()
	var src *os.File
	var err error
	if logPath == "" {
		src = os.Stdout
	} else {
		if err = os.MkdirAll(logPath, 0777); err != nil {
			return nil, err
		}
		logFileName := now.Format("2006-01-02") + ".log"

		fileName := path.Join(logPath, logFileName)
		if _, err = os.Stat(fileName); err != nil {
			if _, err = os.Create(fileName); err != nil {
				return nil, err
			}
		}

		if src, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
			return nil, err
		}
	}

	return &logrus.Logger{
		Out: src,
		Formatter: &logFormatter{
			enableLogColor: enableLogColor,
		},
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: true,
	}, nil
}
