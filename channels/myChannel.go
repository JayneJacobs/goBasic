package channels

import (
	"fmt"
	"time"
)

// DemoChannel creates a channel
func DemoChannel() {
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil")
		a = make(chan int)
		fmt.Printf("This is the channel type %T\n", a)
	}
}

// d := <- channel
// channel <- d

// sends and receives to a channel are blocking
// when data is sent to a channel send is blocked until a go routine reads from the channel.
// when data is read, read is now blocked until data is sent to a channel.
// This helps go routines communicate without locks.

// DemoChannel1 Shows how to add a wait channel
func DemoChannel1(d chan bool) {

	fmt.Println("In Demo Channel1\n")
	time.Sleep(5 * time.Second)
	fmt.Println("In Demo Channel1 Waking from sleep\n")
	d <- true
}

// MainFuncGoRoutine simulating main Go routine
func MainFuncGoRoutine(d chan bool) {

	go DemoChannel1(d)
	fmt.Println("In Main Go Routine calling DemoChannel1\n")
	d <- true
}

// RunGoRoutines runs the simulated Main go ruoutine
func RunGoRoutines() {
	d := make(chan bool)

	go MainFuncGoRoutine(d)

	fmt.Println("In Run Go Routine\n")
	<-d
}
