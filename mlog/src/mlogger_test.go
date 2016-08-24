package mlogger

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
func Test_Printf(t *testing.T) {
	l := New(3)

	for i := 0; i < 200; i++ {
		l.Printf("This is log")
	}

}
*/

const (
	LOGLEVEL = 4
)

func Test_Concurrent_Printf(t *testing.T) {
	fmt.Println(time.Now().Format(TIMEFORMAT))
	l := New(LOGLEVEL)

	for i := 0; i < 200; i++ {
		go l.Fatalf("This is log: %s\n", strconv.Itoa(i))
	}

	fmt.Println(time.Now().Format(TIMEFORMAT))
	time.Sleep(time.Duration(5) * time.Second)
}

/*
func Test_Concurrent_LockPrintf(t *testing.T) {
	fmt.Println(time.Now().Format(TIMEFORMAT))
	l := New(3)
	l.Logf(DEBUG, "abc")
	for i := 0; i < 500; i++ {
		//l.LockPrintf("This is log %s", strconv.Itoa(i))
	}

	l.Logf(INFO, "abc")
	fmt.Println(time.Now().Format(TIMEFORMAT))
	time.Sleep(time.Duration(5) * time.Second)
	l.Logf(TRACE, "abc")
}

func Test_Sequence_LockPrintf(t *testing.T) {
	l := New(LOGLEVEL)
	l.Infof("INFO")
	l.Debugf("DEBUG")
	l.Errorf("ERROR")
	l.Warnf("WARN")
	l.Tracef("TRACE")
	l.Fatalf("FATAL")
}
*/
