package ovo

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// LogLevel is type level of log.
type LogLevel int8

// Available options for LogLevel.
const (
	NoLog LogLevel = iota
	LogError
	LogInfo
	LogDebug
)

// Logger is logging interface.
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type logger struct {
	level LogLevel
}

func defaultLogger(level LogLevel) *logger {
	return &logger{
		level: level,
	}
}

// Debug to print debug log.
func (l *logger) Debug(format string, args ...interface{}) {
	if l.level >= LogDebug {
		fmt.Fprintf(os.Stdout, "[D] "+format+"\n", args...)
	}
}

// Info to print info log.
func (l *logger) Info(format string, args ...interface{}) {
	if l.level >= LogInfo {
		fmt.Fprintf(os.Stdout, "[I] "+format+"\n", args...)
	}
}

// Error to print error log.
func (l *logger) Error(format string, args ...interface{}) {
	if l.level >= LogError {
		_, f, l, _ := runtime.Caller(1)
		caller := filename(f) + ":" + strconv.Itoa(l)
		fmt.Fprintf(os.Stderr, "[E] "+caller+": "+format+"\n", args...)
	}
}

func filename(fpath string) string {
	if i := strings.LastIndexByte(fpath, filepath.Separator); i >= 0 {
		return fpath[i+1:]
	}
	return fpath
}
