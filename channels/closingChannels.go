package channels

import "fmt"

//senders can use a receiver to notify that no more data is on a channel

// A value read from a closed channel will be the null value of the type.

func closingChannel(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

// CloseChannel make and claose a channel
func CloseChannel() {
	channel := make(chan int)
	go closingChannel(channel)
	for {
		x, y := <-channel
		if y == false {
			break
		}
		fmt.Println("Values and result", x, y)
	}
}

// RangeChannel Ranges over teh items until channel closes
func RangeChannel() {
	channel := make(chan int)

	go closingChannel(channel)

	for i := range channel {
		fmt.Println("This is the result of Range", i)
	}
}
