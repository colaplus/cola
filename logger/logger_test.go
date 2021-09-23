package logger

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "d:/logcenter/", "test")
	logger.Debug("user id[%d] is come from china\n", 2342343)
	logger.Warn("test warn log\n")
	logger.Fatal("test Fatal log\n")
	logger.Close()
}

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "d:/logcenter/", "test")
	logger.Debug("user id[%d] is come from china\n", 2342343)
	logger.Warn("test warn log\n")
	logger.Fatal("test Fatal log\n")
	logger.Close()
}
