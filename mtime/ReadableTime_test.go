package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_Sub(t *testing.T) {
	start := time.Now()
	time.Sleep(time.Duration(5) * time.Second)
	now := time.Now()

	mt := &MyTime{
		Start: start.Format(time.RFC3339),
		Now:   now.Format(time.RFC3339),
	}

	mt.Sub()

	fmt.Printf("Start: %s", mt.Start)
	fmt.Printf("now: %s", mt.Now)
	fmt.Printf("Duration: %s", mt.Duration)

	mt.Localize()
	fmt.Printf("Start: %s", mt.Start)
	fmt.Printf("now: %s", mt.Now)
	fmt.Printf("Duration: %s", mt.Duration)

}
