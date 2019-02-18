package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

// Sleeper is an interface for a sleep function
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is the sleeper used in this package for sleeping
type DefaultSleeper struct{}

// Sleep will sleep
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown counts down
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
