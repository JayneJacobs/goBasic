package defering

import "fmt"

func deferedMethod(y int) {
	fmt.Println("y = ", y)
}

// RunDeferedMethod shows which is read first
func RunDeferedMethod() {
	xy := 10
	defer deferedMethod(xy)
	xy = 20 // this is read first
	fmt.Println("xy = ", xy)
}
