package main

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	format := "%{color}Message -> %{message}"
	logCustom, _ := GetLoggerFormatted("log_custom", format)
	logCustom.Critical("CUSTOM")
	logDefault, _ := GetLogger("log_default")
	logDefault.Critical("DEFAULT")
	logCustom.Error("cust2")
	
}

