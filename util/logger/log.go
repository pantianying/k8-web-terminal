package logger

import (
	"github.com/astaxie/beego/logs"
)

var (
	l *logs.BeeLogger
)

func init() {
	l := logs.NewLogger()
	l.SetLogger(logs.AdapterFile,
		`{"filename":"k8-web-terminal.log","maxlines":0,"maxsize":0,
"daily":true,"maxdays":10,"color":true}`)
	l.SetLogger(logs.AdapterConsole)
	l.SetLevel(logs.LevelInfo)
}

func Debug(format string, v ...interface{}) {
	l.Debug(format, v)
}

func Info(format string, v ...interface{}) {
	l.Info(format, v)
}

func Error(format string, v ...interface{}) {
	l.Error(format, v)
}
