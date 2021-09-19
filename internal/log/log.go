package log

import (
	"os"

	"github.com/op/go-logging"
)

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}[%{time:15:04:05.000}] %{longfile} %{longfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// GetLogger new logger
func GetLogger(name string) *logging.Logger {
	if name == "" {
		name = "server"
	}
	stdoutBE := logging.NewLogBackend(os.Stdout, "", 0)
	beFormat := logging.NewBackendFormatter(stdoutBE, format)
	logging.SetBackend(beFormat)
	var log1 = logging.MustGetLogger(name)
	return log1
}

// Default logger
var Default = GetLogger("server")

// Error logs a message using ERROR as log level.
var Error = Default.Error

// Errorf logs a message using ERROR as log level.
var Errorf = Default.Errorf

// Info logs a message using INFO as log level.
var Info = Default.Info

// Infof logs a message using INFO as log level.
var Infof = Default.Infof

// Debug logs a message using DEBUG as log level.
var Debug = Default.Debug

// Debugf logs a message using DEBUG as log level.
var Debugf = Default.Debugf

// Critical logs a message using CRITICAL as log level.
var Critical = Default.Critical

// Criticalf logs a message using CRITICAL as log level.
var Criticalf = Default.Criticalf

// Warning logs a message using WARNING as log level.
var Warning = Default.Warning

// Warningf logs a message using WARNING as log level.
var Warningf = Default.Warningf

// Notice logs a message using NOTICE as log level.
var Notice = Default.Notice

// Noticef logs a message using NOTICE as log level.
var Noticef = Default.Noticef

// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
var Panic = Default.Panic

// Panicf is equivalent to l.Critical followed by a call to panic().
var Panicf = Default.Panicf

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
var Fatal = Default.Fatal

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
var Fatalf = Default.Fatalf
