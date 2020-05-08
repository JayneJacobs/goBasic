package functionhigherorder

import (
	"fmt"
)

//functions that take functions as arguments

func s(x func(xy int) int) {
	fmt.Println(x(222))
}

// DemoHiOrderFunc passes a func as an arguement to a function
func DemoHiOrderFunc() {
	z := func(xy int) int {
		return xy
	}
	s(z)
}
