package clog

import (
	"fmt"
	"log"
)

type TextColor string

var Black TextColor = "[30m"
var Red TextColor = "[31m"
var Green TextColor = "[32m"
var Yellow TextColor = "[33m"
var Blue TextColor = "[34m"
var Magenta TextColor = "[35m"
var Cyan TextColor = "[36m"
var White TextColor = "[37m"
var BrightBlack TextColor = "[90m"
var BrightRed TextColor = "[91m"
var BrightGreen TextColor = "[92m"
var BrightYellow TextColor = "[93m"
var BrightBlue TextColor = "[94m"
var BrightMagenta TextColor = "[95m"
var BrightCyan TextColor = "[96m"
var BrightWhite TextColor = "[97m"

var escape = "\x1b"
var defaults = fmt.Sprintf("%s[0m", escape)

func Colorize(text string, color TextColor) string {
	return fmt.Sprintf("%s%s%s%s", escape, color, text, defaults)
}

type LogLevel int

const (
	Debug   LogLevel = iota
	Info    LogLevel = iota
	Warning LogLevel = iota
	Error   LogLevel = iota
)

type Logger struct {
	logLevel LogLevel
}

func NewLogger(logLevel LogLevel) Logger {
	return Logger{logLevel: logLevel}
}

func (l *Logger) Debug(msg string) {
	if l.logLevel > Debug {
		return
	}
	log.Printf("DEBUG:   %s", Colorize(msg, BrightBlue))
}

func (l *Logger) Info(msg string) {
	if l.logLevel > Info {
		return
	}
	log.Printf("INFO:    %s", Colorize(msg, BrightGreen))
}

func (l *Logger) Warning(msg string) {
	if l.logLevel > Warning {
		return
	}
	log.Printf("WARNING: %s", Colorize(msg, BrightYellow))
}

func (l *Logger) Error(msg string) {
	log.Printf("ERROR:   %s", Colorize(msg, BrightRed))
}
