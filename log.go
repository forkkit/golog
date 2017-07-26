package golog

import (
	"reflect"
	"sync"
	"time"
)

var zeroTime = reflect.Zero(reflect.TypeOf(time.Time{})).Interface()

// A Log represents a log line.
type Log struct {
	timeFormat string
	Time       time.Time
	Level      Level
	Message    string
	// NewLine returns false if this Log
	// derives from a `Print` function,
	// otherwise true if derives from a `Println`, `Error`, `Errorf`, `Warn`, etc...
	//
	// This NewLine does not mean that `Message` ends with "\n".
	// NewLine has to do with the methods called,
	// not the original content of the `Message`.
	NewLine bool
}

// FormatTime returns the formatted `Time`.
func (l Log) FormatTime() string {
	return l.Time.Format(l.timeFormat)
}

var logPool = sync.Pool{New: func() interface{} {
	return Log{}
}}

// acquireLog returns a new log fom the pool.
func acquireLog(level Level, msg string, withPrintln bool, timeFormat string) Log {
	l := logPool.Get().(Log)
	l.NewLine = withPrintln
	l.timeFormat = timeFormat
	l.Time = time.Now()
	l.Level = level
	l.Message = msg
	return l
}

// releaseLog Log releases a log instance back to the pool.
func releaseLog(l Log) {
	logPool.Put(l)
}
