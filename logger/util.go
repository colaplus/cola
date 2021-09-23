package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogData struct {
	Message      string
	TimeStr      string
	LevelStr     string
	FileName     string
	FuncName     string
	LineNo       int
	WarnAndFatal bool
}

func getLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}

	return
}

func writelog(level int, format string, args ...interface{}) *LogData {
	now := time.Now()
	nowStr := now.Format("2006-1-2 15:04:05 ")
	levelStr := getLevelText(level)
	fileName, funcName, lineNo := getLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)

	msg := fmt.Sprintf(format, args...)

	logData := &LogData{
		Message:  msg,
		TimeStr:  nowStr,
		LevelStr: levelStr,
		FileName: fileName,
		FuncName: funcName,
		LineNo:   lineNo,
		WarnAndFatal: false,
	}

	if level== LogLevelError || level==LogLevelWarn || level ==LogLevelFatal {
		logData.WarnAndFatal = true
	}

	return logData
	// fmt.Fprintf(file, "%s %s %s:%d %s %s\n", nowStr, levelStr, fileName, lineNo, funcName, msg)
}
