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
	FATAL: "Fatal",
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
	Level  LogLevel
	lock   *sync.Mutex
	rwlock *sync.RWMutex
}

func New(level int) *Mlogger {
	return &Mlogger{Level: LogLevel(level), lock: new(sync.Mutex), rwlock: new(sync.RWMutex)}
}

func (m *Mlogger) Printf(format string, args ...interface{}) {
	time := time.Now().Format(TIMEFORMAT)
	loglevel := m.Level.Level()
	pc, file, line, _ := runtime.Caller(CALLERDEPTH)
	callFunc := runtime.FuncForPC(pc).Name()
	context := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s [%s] %s line %d: %s: %s\n", time, loglevel, file, line, callFunc, context)
}

func (m *Mlogger) LockPrintf(level LogLevel, format string, args ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	time := time.Now().Format(TIMEFORMAT)
	loglevel := level.Level()
	pc, file, line, _ := runtime.Caller(CALLERDEPTH)
	callFunc := runtime.FuncForPC(pc).Name()
	context := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s [%s] %s line %d: %s: %s\n", time, loglevel, file, line, callFunc, context)
}

func (m *Mlogger) Logf(level LogLevel, format string, args ...interface{}) {
	if m.Level >= FATAL {
		m.LockPrintf(FATAL, format, args...)
	} else if m.Level >= ERROR {
		m.LockPrintf(ERROR, format, args...)
	} else if m.Level >= WARN {
		m.LockPrintf(WARN, format, args...)
	} else if m.Level >= INFO {
		m.LockPrintf(INFO, format, args...)
	} else if m.Level >= DEBUG {
		m.LockPrintf(DEBUG, format, args...)
	} else if m.Level >= TRACE {
		m.LockPrintf(TRACE, format, args...)
	}
}

func (m *Mlogger) Fatalf(format string, args ...interface{}) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	if m.Level >= FATAL {
		m.LockPrintf(FATAL, format, args...)
	}
}

func (m *Mlogger) Errorf(format string, args ...interface{}) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	if m.Level >= ERROR {
		m.LockPrintf(ERROR, format, args...)
	}
}

func (m *Mlogger) Warnf(format string, args ...interface{}) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	if m.Level >= ERROR {
		m.LockPrintf(WARN, format, args...)
	}
}

func (m *Mlogger) Infof(format string, args ...interface{}) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	if m.Level >= INFO {
		m.LockPrintf(INFO, format, args...)
	}
}

func (m *Mlogger) Debugf(format string, args ...interface{}) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	if m.Level >= DEBUG {
		m.LockPrintf(DEBUG, format, args...)
	}
}

func (m *Mlogger) Tracef(format string, args ...interface{}) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	if m.Level >= TRACE {
		m.LockPrintf(TRACE, format, args...)
	}
}
