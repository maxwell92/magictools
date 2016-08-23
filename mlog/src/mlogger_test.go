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

	for i := 0; i < 200; i++ {
		go l.Printf("This is log %s", strconv.Itoa(i))
	}

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println(time.Now().Format(TIMEFORMAT))
}

func Test_Concurrent_LockPrintf(t *testing.T) {
	fmt.Println(time.Now().Format(TIMEFORMAT))
	l := New(3)

	for i := 0; i < 200; i++ {
		go l.Printf("This is log %s", strconv.Itoa(i))
	}

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println(time.Now().Format(TIMEFORMAT))
}
