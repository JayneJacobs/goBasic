package chanselect

import (
	"fmt"
	"time"
)

func chan1(channel chan string) {
	time.Sleep(6 * time.Second)
	channel <- "from chan1"
}

func chan2(channel chan string) {
	time.Sleep(3 * time.Second)
	channel <- "from chan2"
}

func RunChannelSelect() {

	out1 := make(chan string)
	out2 := make(chan string)
	go chan2(out2)
	go chan1(out1)

	select {
	case a1 := <-out1:
		fmt.Println(a1)
	case a2 := <-out2:
		fmt.Println(a2)
		// default:
		// fmt.Println("No value received") //always runs first
	} // for load dependant cases. Server that responds fastest is heard
}
