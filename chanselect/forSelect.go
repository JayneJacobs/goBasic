package chanselect

import (
	"fmt"
	"time"
)

func case1(channel chan string) {
	time.Sleep(1050 * time.Millisecond)
	channel <- "case1 data entered"
}

// RunforSelect shows 2 cases with default
func RunforSelect() {
	channel := make(chan string)
	go case1(channel)
	for {
		time.Sleep(1000 * time.Microsecond)
		select {
		case vx := <-channel:
			fmt.Println("Received value: ", vx)
			return
		default:
			fmt.Println("Value not received")
		}
	}
}

// RunOnlyDefault demos the default case role
func RunOnlyDefault() {
	channel := make(chan string)
	select {
	case <-channel:
	default:
		fmt.Println("Default Case") //prevents block
	}
}

// TwoCasesWDefault executes without blocking
func TwoCasesWDefault() {
	var channel chan string
	select {
	case vx := <-channel:
		fmt.Println("value", vx)
	default:
		fmt.Println("Default case")
	}

}

// EmptySelect shows foreveer blodk and Panic
func EmptySelect() {
	select {}
}
