package mlogger

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	TIMEFORMAT  = "2006-01-02 15:04:05.0000"
	CALLERDEPTH = 2
)

type LogLevel uintptr

const (
	FATAL LogLevel = 1 /*Fatal Error*/
	ERROR LogLevel = 2 /*Error*/
	WARN  LogLevel = 3 /*Warning*/
	INFO  LogLevel = 4 /*Information*/
	DEBUG LogLevel = 5 /*Debug*/
	TRACE LogLevel = 6 /*Trace*/
)

var loglevels = [...]string{
	FATAL: "Fatal Error",
	ERROR: "Error",
	WARN:  "Warning",
	INFO:  "Info",
	DEBUG: "Debug",
	TRACE: "Trace",
}

func (l LogLevel) Level() string {
	return loglevels[l]
}

type Mlogger struct {
	Level LogLevel
	lock  *sync.Mutex
}

func New(level int) *Mlogger {
	return &Mlogger{Level: LogLevel(level), lock: new(sync.Mutex)}
}

func (m *Mlogger) Printf(v ...interface{}) {
	time := time.Now().Format(TIMEFORMAT)
	loglevel := m.Level.Level()
	_, file, _, _ := runtime.Caller(CALLERDEPTH)
	_, _, line, _ := runtime.Caller(CALLERDEPTH)
	context := fmt.Sprintf("%s", v)
	fmt.Fprintf(os.Stderr, "%s [%s] %s line %d: %s\n", time, loglevel, file, line, context)
}

func (m *Mlogger) LockPrintf(v ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	time := time.Now().Format(TIMEFORMAT)
	loglevel := m.Level.Level()
	_, file, _, _ := runtime.Caller(CALLERDEPTH)
	_, _, line, _ := runtime.Caller(CALLERDEPTH)
	context := fmt.Sprintf("%s", v)
	fmt.Fprintf(os.Stderr, "%s [%s] %s line %d: %s\n", time, loglevel, file, line, context)
}

/*
func (m *Mlogger) Println() {
	fmt.Fprintln(os.Stderr, "abc")
}
*/
