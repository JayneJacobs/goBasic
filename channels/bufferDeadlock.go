package channels

import "fmt"

// BeyondCapacity will Panic
func BeyondCapacity() {
	channel := make(chan string, 2)
	channel <- "ABC"
	channel <- "DEF"
	channel <- "This is outside of cap"

	fmt.Println(<-channel)

	fmt.Println(<-channel)
}

// BeyondLength shows that len varies
func BeyondLength() {
	channel := make(chan string, 3)
	channel <- "ABC"
	channel <- "DEF"
	channel <- "This is in capacity"

	fmt.Println("channel capacity", cap(channel))
	fmt.Println("channel capacity", len(channel))

	fmt.Println("read value: ", <-channel)
	fmt.Println("channel capacity", len(channel))
}
