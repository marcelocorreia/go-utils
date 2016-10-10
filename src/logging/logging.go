package logging

import (
	"github.com/op/go-logging"
	"os"
	"strings"
)

type LoggingError struct {
	S string
}

func (e *LoggingError) Error() string {
	return e.S
}

const DEFAULT_FORMAT = "%{color}%{time:20060201 15:04:05.000} %{shortfunc} %{level:.4s} %{id:03x}%{color:reset} %{message}"

func GetLogger(loggerName string)(*logging.Logger, error){
	return GetLoggerFormatted(loggerName,nil)
}

func GetLoggerFormatted(loggerName string, stringFormatter interface{}) (*logging.Logger, error)  {
	if loggerName == "" {
		newError := &LoggingError{"logger name needs to be defined"}
		return nil, newError
	}
	
	logger := logging.MustGetLogger(loggerName)
	var format logging.Formatter
	if stringFormatter != nil {
		format = logging.MustStringFormatter(stringFormatter.(string))
	} else {
		format = logging.MustStringFormatter(DEFAULT_FORMAT)
	}
	StdOutBackend := logging.NewLogBackend(os.Stdout, "", 0)
	StdOutBackendLeveled := logging.AddModuleLevel(StdOutBackend)

	StdErrBackend := logging.NewLogBackend(os.Stderr, "", 0)
	StdErrBackendLeveled := logging.AddModuleLevel(StdErrBackend)

	
	level := strings.ToLower(os.Getenv(strings.ToUpper(loggerName) + "_LOG_LEVEL"))
	switch level {
	case "error":
		StdOutBackendLeveled.SetLevel(logging.ERROR, "")
		StdErrBackendLeveled.SetLevel(logging.ERROR, "")
	case "warning":
		StdOutBackendLeveled.SetLevel(logging.WARNING, "")
		StdErrBackendLeveled.SetLevel(logging.WARNING, "")
	case "info":
		StdOutBackendLeveled.SetLevel(logging.INFO, "")
		StdErrBackendLeveled.SetLevel(logging.INFO, "")
	case "critical":
		StdOutBackendLeveled.SetLevel(logging.CRITICAL, "")
		StdErrBackendLeveled.SetLevel(logging.CRITICAL, "")
	case "debug":
		StdOutBackendLeveled.SetLevel(logging.DEBUG, "")
		StdErrBackendLeveled.SetLevel(logging.DEBUG, "")
	case "notice":
		StdOutBackendLeveled.SetLevel(logging.NOTICE, "")
		StdErrBackendLeveled.SetLevel(logging.NOTICE, "")
	default:
		StdOutBackendLeveled.SetLevel(logging.ERROR, "")
		StdErrBackendLeveled.SetLevel(logging.ERROR, "")
	}
	
	logging.SetBackend(StdOutBackendLeveled)
	logging.SetFormatter(format)
	
	return logger, nil
}