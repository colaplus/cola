package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileLogger struct {
	level        int
	logPath      string
	logName      string
	file         *os.File
	warnFile     *os.File
	LogDataChan  chan *LogData
	logSplitType int
	logSplitSize int64
	lastSplitHour int
}

func NewFileLogger(config map[string]string) (logger LogInterface, err error) {
	logPath, ok := config["logPath"]
	if !ok {
		err = fmt.Errorf("not found logPath")
		return
	}
	logName, ok := config["logName"]
	if !ok {
		err = fmt.Errorf("not found logName")
		return
	}
	logLevel, ok := config["logLevel"]
	if !ok {
		err = fmt.Errorf("not found logLevel")
		return
	}
	logChanSize, ok := config["LogChanSize"]
	if !ok {
		logChanSize = "50000"
	}
	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}

	var logSplitType int = LogSplitTypeHour
	var logSplitSize int64
	logSplitTypeStr, ok := config["logSplitTypeStr"]
	if !ok {
		logSplitTypeStr = "hour"
	} else {
		if logSplitTypeStr == "size" {
			logSplitSizeStr, ok := config["logSplitSize"]
			if !ok {
				logSplitSizeStr = "104857600"
			}

			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 104857600
			}

			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}
	logger = &FileLogger{
		level:       getLogLevel(logLevel),
		logPath:     logPath,
		logName:     logName,
		LogDataChan: make(chan *LogData, chanSize),
		logSplitType: logSplitType,
		logSplitSize: logSplitSize,
	}
	logger.init()

	return logger, err
}

func (f *FileLogger) checkSplitFile(warnFile bool) {
	if f.logSplitType == LogSplitTypeHour {
		now := time.Now()
		hour := now.Hour()
		if hour == f.lastSplitHour {
			return
		}

		var backupFileName string
		var fileName string
		if warnFile {
			backupFileName = fmt.Sprintf("%s%s.log.wf_%04d%02d%02d%02d", f.logPath, f.logName, now.Year(),now.Month(), now.Day(), f.lastSplitHour)
			fileName = fmt.Sprintf("%s%s.log.wf", f.logPath, f.logName)
		} else {
			backupFileName = fmt.Sprintf("%s%s.log_%04d%02d%02d%02d", f.logPath, f.logName, now.Year(),now.Month(), now.Day(), f.lastSplitHour)
			fileName = fmt.Sprintf("%s%s.log", f.logPath, f.logName)

		}

		file := f.file
		if warnFile {
			file = f.warnFile
		}

		file.Close()
		os.Rename(fileName, backupFileName)

		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			return
		}

		if warnFile {
			f.warnFile = file
		} else {
			f.file = file
		}
	}
}

func (f *FileLogger) init() {
	fileName := fmt.Sprintf("%s%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", fileName, err))
	}

	f.file = file

	// 写错误日志和fatal日志的文件
	fileName = fmt.Sprintf("%s%s.wf.log", f.logPath, f.logName)
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", fileName, err))
	}

	f.warnFile = file
	go f.writeLogBackground()
}

func (f *FileLogger) writeLogBackground() {
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnFile
		}

		f.checkSplitFile(logData.WarnAndFatal)
		fmt.Fprintf(file, "%s %s %s:%d %s %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Message)
	}
}

func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	logData := writelog(LogLevelDebug, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writelog(LogLevelTrace, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writelog(LogLevelInfo, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := writelog(LogLevelWarn, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := writelog(LogLevelError, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := writelog(LogLevelFatal, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
