package main

// Microservice logging

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/term"
	"github.com/go-stack/stack"
)

// ServiceLogger is wrapper for go-kit custom logger specifically giving colour & extra info
type ServiceLogger interface {
	With(keyvals ...interface{}) ServiceLogger
	WithPrefix(keyvals ...interface{}) ServiceLogger

	Debug(keyvals ...interface{}) error
	Info(keyvals ...interface{}) error
	Stage(keyvals ...interface{}) error
	Warn(keyvals ...interface{}) error
	Err(keyvals ...interface{}) error
	Crit(keyvals ...interface{}) error

	DebugMsg(msg string) error
	InfoMsg(msg string) error
	StageMsg(msg string) error
	WarnMsg(msg string) error
	ErrMsg(msg string) error
	CritMsg(msg string) error
}

// NewServiceLogger returns a basic ServiceLogger with all of the expected functionality
func NewServiceLogger(serviceName string) ServiceLogger {

	// Color by level value
	var colorFn = func(keyvals ...interface{}) term.FgBgColor {
		for i := 0; i < len(keyvals)-1; i += 2 {
			if keyvals[i] != "level" {
				continue
			}
			switch keyvals[i+1] {
			case "debug":
				return term.FgBgColor{Fg: term.DarkGray, Bg: term.Default}
			case "info":
				return term.FgBgColor{Fg: term.DarkGreen, Bg: term.Default}
			case "stage":
				return term.FgBgColor{Fg: term.DarkBlue, Bg: term.Default}
			case "warn":
				return term.FgBgColor{Fg: term.Yellow, Bg: term.Default}
			case "error":
				return term.FgBgColor{Fg: term.Red, Bg: term.Default}
			case "crit":
				return term.FgBgColor{Fg: term.Gray, Bg: term.DarkRed}
			default:
				return term.FgBgColor{Fg: term.Default, Bg: term.Default}
			}
		}
		return term.FgBgColor{Fg: term.Default, Bg: term.Default}
	}

	var ts = func() log.Valuer {
		return func() interface{} {
			return time.Now().Format("02/01/2006 03:04:05")
		}

	}

	var kitLogger log.Logger
	{
		kitLogger = term.NewColorLogger(os.Stdout, log.NewLogfmtLogger, colorFn)
		kitLogger = log.With(kitLogger, "ts", ts())
		kitLogger = log.With(kitLogger, "service", serviceName)
	}

	logger := BasicServiceLogger{kitLogger}

	return logger
}

// -----------------------------------------------------------------------------

// BasicServiceLogger returns a naïve, stateless implementation of ServiceLogger.
type BasicServiceLogger struct {
	log.Logger
}

// -----------------------------------------------------------------------------

//With prints a log message with a keyval
func (s BasicServiceLogger) With(keyvals ...interface{}) ServiceLogger {
	logger := log.With(s.Logger, keyvals...)
	return BasicServiceLogger{logger}
}

//WithPrefix prints a log message with a prefix
func (s BasicServiceLogger) WithPrefix(keyvals ...interface{}) ServiceLogger {
	logger := log.WithPrefix(s.Logger, keyvals...)
	return BasicServiceLogger{logger}
}

//------------------------------------------------------------------------------
// Log levels functions

//Debug prints a log message with debug level
func (s BasicServiceLogger) Debug(keyvals ...interface{}) error {
	s.Logger = log.With(s, "level", "debug", "caller", stack.Caller(2))
	err := s.Log(keyvals...)
	return err
}

//Info prints a log message with info level
func (s BasicServiceLogger) Info(keyvals ...interface{}) error {
	s.Logger = log.With(s, "level", "info", "caller", stack.Caller(2))
	err := s.Log(keyvals...)
	return err
}

//Stage prints a log message with stage level
func (s BasicServiceLogger) Stage(keyvals ...interface{}) error {
	s.Logger = log.With(s, "level", "stage", "caller", stack.Caller(2))
	err := s.Log(keyvals...)
	return err
}

//Warn prints a log message with warn level
func (s BasicServiceLogger) Warn(keyvals ...interface{}) error {
	s.Logger = log.With(s, "level", "warn", "caller", stack.Caller(2))
	err := s.Log(keyvals...)
	return err
}

//Err prints a log message with err level
func (s BasicServiceLogger) Err(keyvals ...interface{}) error {
	s.Logger = log.With(s, "level", "error", "caller", stack.Caller(2))
	err := s.Log(keyvals...)
	return err
}

//Crit prints a log message with crit level
func (s BasicServiceLogger) Crit(keyvals ...interface{}) error {
	s.Logger = log.With(s, "level", "crit", "caller", stack.Caller(2))
	err := s.Log(keyvals...)
	println("")
	return err
}

// -----------------------------------------------------------------------------
// Simplified msg function (only one string ie. no need for "msg", <val>)

//DebugMsg prints a log message with debug level
func (s BasicServiceLogger) DebugMsg(msg string) error {
	s.Logger = log.With(s, "level", "debug", "caller", stack.Caller(2))
	err := s.Log("msg", msg)
	return err
}

//InfoMsg prints a log message with info level
func (s BasicServiceLogger) InfoMsg(msg string) error {
	s.Logger = log.With(s, "level", "info", "caller", stack.Caller(2))
	err := s.Log("msg", msg)
	return err
}

//StageMsg prints a log message with stage level
func (s BasicServiceLogger) StageMsg(msg string) error {
	s.Logger = log.With(s, "level", "stage", "caller", stack.Caller(2))
	err := s.Log("msg", msg)
	return err
}

//WarnMsg prints a log message with warn level
func (s BasicServiceLogger) WarnMsg(msg string) error {
	s.Logger = log.With(s, "level", "warn", "caller", stack.Caller(2))
	err := s.Log("msg", msg)
	return err
}

//ErrMsg prints a log message with err level
func (s BasicServiceLogger) ErrMsg(msg string) error {
	s.Logger = log.With(s, "level", "error", "caller", stack.Caller(2))
	err := s.Log("msg", msg)
	return err
}

//CritMsg prints a log message with crit level
func (s BasicServiceLogger) CritMsg(msg string) error {
	s.Logger = log.With(s, "level", "crit", "caller", stack.Caller(2))
	err := s.Log("msg", msg)
	println("")
	return err
}
