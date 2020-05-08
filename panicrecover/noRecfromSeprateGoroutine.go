package panicrecover

import (
	"fmt"
	"time"
)

// You cannot use recover in a separate go routine

func recoverfunc3() {
	if rx := recover(); rx != nil {
		fmt.Println("recovered from ", rx)
	}
}

func a() {
	defer recoverfunc3() //cannot recover in x from a panic in y
	fmt.Println("in func X")
	b() // not a go routine
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("in func Y")
	panic("func y panicked") //there is a panick in y.
}

// RunDemonoRec demonstrates
func RunDemoRec() {
	a()
	fmt.Println("back to main")
}
