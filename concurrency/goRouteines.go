package concurrency

import (
	"fmt"
	"time"
)

//Concurrency is not Parralellism Currency is sequential scheduling of go routines on a single Thread
//Threads are Parrallel.
//Parralellism is not faster if process wait on each other to communicate.
//Concurrency results in less resource intensive communication
//Go routines are functions / methods that run concurrently with other methods.
//Go Routines act like lightweight threads. Go provides the signallling  They are cheaper than threads. Thousands of Go Routines per thread.

// Channels prevent race conditions.

// RoutineOne is called as a go Routine
func RoutineOne() {
	fmt.Println("HI this is a Go Routine. I had time to Execute!!!\n")
}

// DemoConcurrency shows that a race condition can occur with simultaneous routines.
func DemoConcurrency() {
	go RoutineOne()
	fmt.Println("This is after teh Go Routine, Main Exits before runing the go routine\n")
}

func DemoAddTiming() {

	time.Sleep(1 * time.Second)
	fmt.Println("This is after teh Go Routine, Sleep before the Main Function Executes gives the go routine time to complete\n")
}

func DemoAddMoreTiming() {
	go RoutineOne()
	time.Sleep(2 * time.Second)
	fmt.Println("This is after teh Go Routine, Sleep before the Main Function Executes gives the go routine time to complete\n")
}
