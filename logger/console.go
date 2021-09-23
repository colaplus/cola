package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger(config map[string]string) (logger LogInterface, err error) {
	logLevel, ok := config["logLevel"]
	if !ok {
		err = fmt.Errorf("not found loglevel")
	}

	logger = &ConsoleLogger{
		level: getLogLevel(logLevel),
	}

	return logger, err
}

func (f *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *ConsoleLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	logData := writelog(LogLevelDebug, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Message)
}
func (f *ConsoleLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writelog(LogLevelTrace, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Message)
}
func (f *ConsoleLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writelog(LogLevelInfo, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Message)
}
func (f *ConsoleLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := writelog(LogLevelWarn, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Message)
}
func (f *ConsoleLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := writelog(LogLevelError, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Message)
}
func (f *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := writelog(LogLevelFatal, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Message)
}
func (f *ConsoleLogger) Close() {}
func (f *ConsoleLogger) init()  {}
