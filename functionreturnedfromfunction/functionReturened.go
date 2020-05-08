package functionreturnedfromfunction

import "fmt"

//functions that take functions as arguments

func s() func(xy int) int {
	z := func(xy int) int {
		return xy
	}
	return z
}

// DemoHiOrderFunc passes a func as an arguement to a function
func DemoHiOrderFunc() {
	g := s()
	fmt.Println(g(333))

}
