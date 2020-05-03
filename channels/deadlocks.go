package channels

import (
	"fmt"
)

// if a go routine is receiving data on a channel it is expected that some thing is writing to the channel. If not the program
// will panic with a deadlock.

func MainLike() {
	a := make(chan int)
	a <- 100
}

func InChannel(xa chan<- int) {
	xa <- 10
}

func MainLikeIn() {
	// xa := make(chan int) //chan<-int not right for one way
	// go InChannel(xa)
	// fmt.Println("InMainlikeIn")
	// fmt.Println(<-xa)
}

func RunMainLike() {
	// go MainLike()
	xa := make(chan int) //chan<-int not right for one way
	go InChannel(xa)
	go MainLikeIn()
	fmt.Println("In RunMainLike")
	fmt.Println(<-xa)
}
