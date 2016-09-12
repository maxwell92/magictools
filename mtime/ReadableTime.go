package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type MyTime struct {
	Start    string `json:"start"`
	Now      string `json:"now"`
	Duration string `json:"duration"`
}

func (t *MyTime) UTC() {
	start, _ := time.Parse(time.RFC3339, t.Start)
	t.Start = start.UTC().Format(time.RFC3339)

	now, _ := time.Parse(time.RFC3339, t.Now)
	t.Now = now.UTC().Format(time.RFC3339)

}

func (t *MyTime) Localize() {
	start, _ := time.Parse(time.RFC3339, t.Start)
	t.Start = start.Local().Format(time.RFC3339)

	now, _ := time.Parse(time.RFC3339, t.Now)
	t.Now = now.Local().Format(time.RFC3339)
}

func (t *MyTime) Sub() {

	zstart, _ := time.Parse(time.RFC3339, t.Start)
	start := zstart.UTC()

	znow := time.Now()

	now := znow.UTC()
	t.Now = now.Format(time.RFC3339)

	duration := now.Sub(start)
	hour := duration.Hours()
	min := duration.Minutes()
	sec := duration.Seconds()

	if hour > 24 {
		t.Duration = strconv.Itoa(int(hour/24)) + " days"
	} else if hour < 1 {
		if min > 1 {
			t.Duration = strconv.Itoa(int(min)) + " mins"
		} else {
			t.Duration = strconv.Itoa(int(sec)) + " secs"
		}
	} else {
		t.Duration = strconv.Itoa(int(hour)) + " hours"
	}

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Insufficient Arguements")
		os.Exit(1)
	}

	mtime := &MyTime{
		Start: os.Args[1],
	}

	fmt.Printf("Start: %s\n", mtime.Start)
	mtime.Sub()
	fmt.Printf("Now: %s\n", mtime.Now)
	fmt.Printf("Duration: %s\n", mtime.Duration)

	mtime.Localize()
	fmt.Printf("Start: %s\n", mtime.Start)
	fmt.Printf("now: %s\n", mtime.Now)
	fmt.Printf("Duration: %s\n", mtime.Duration)

}
