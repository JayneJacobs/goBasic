package panicrecover

import (
	"fmt"
	"time"
)

// You cannot use recover in a separate go routine

func recoverfunc2() {
	if rx := recover(); rx != nil {
		fmt.Println("recovered from ", rx)
	}
}

func x() {
	defer recoverfunc2() //cannot recover in x from a panic in y
	fmt.Println("in func X")
	go y()
	time.Sleep(1 * time.Second)
}

func y() {
	fmt.Println("in func Y")
	panic("func y panicked") //there is a panick in y.
}

// RunDemonoRec demonstrates
func RunDemonoRec() {
	x()
	fmt.Println("back to main")
}
