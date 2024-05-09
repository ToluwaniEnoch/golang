package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord = "Go!"
	countDownValue = 3
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}


func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownValue; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()

	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}