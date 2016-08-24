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

func Test_Concurrent_Printf(t *testing.T) {
	fmt.Println(time.Now().Format(TIMEFORMAT))
	l := New(3)

	for i := 0; i < 500; i++ {
		go l.Printf("This is log %s", strconv.Itoa(i))
	}

	fmt.Println(time.Now().Format(TIMEFORMAT))
	time.Sleep(time.Duration(5) * time.Second)
}

func Test_Concurrent_LockPrintf(t *testing.T) {
	fmt.Println(time.Now().Format(TIMEFORMAT))
	l := New(3)

	for i := 0; i < 500; i++ {
		go l.LockPrintf("This is log %s", strconv.Itoa(i))
	}

	fmt.Println(time.Now().Format(TIMEFORMAT))
	time.Sleep(time.Duration(5) * time.Second)
}
