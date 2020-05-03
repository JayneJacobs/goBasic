package channels

import (
	"fmt"
	"time"
)

// sends and receives are blocking. It is possible to send channel data to a buffer so a channel will only block when the buffer is full
// a go routine receives only when the channel is full, or blocking.
//size parameter specifies a buffer make(chan, type, capacity)

func BufferedChannel() {
	channel := make(chan string, 2)
	channel <- "ABCD"
	channel <- "EFGHI"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func x(channel chan int) {
	for i := 0; i < 5; i++ {
		channel <- i
		fmt.Println(i, "written to channel")
	}
	close(channel)
}

func RunBufferedMain() {
	channel := make(chan int, 2)
	go x(channel)
	time.Sleep(1 * time.Second)
	for val := range channel {
		fmt.Println("valuee written by x routine is: ", val)
		time.Sleep(2 * time.Second)
	}
}
