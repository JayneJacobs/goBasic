package defering

import "fmt"

func x() {
	fmt.Println("inthe fuction x")
}

//Prints the largest number of the slice
func y(n []int) {
	defer x()
	fmt.Println("executing y")
	m := n[0]
	for _, vx := range n {
		if vx > m {
			m = vx
		}
	}
	fmt.Println(n, "has", m, "as the miggest number")
}

func Rundeferring() {
	n := []int{6, 5, 4, 3, 2, 1}
	y(n)
}
