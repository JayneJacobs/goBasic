package defering

import "fmt"

// Arguements are read LIFO order

// RunLIFOorder shows that for statements take filo order
func RunLIFOorder() {
	nx := "Jayne"
	fmt.Printf("Before: nx = %s \n", string(nx))
	fmt.Printf("######\n")
	for _, vx := range []rune(nx) {
		defer fmt.Printf("%c", vx)
	}
}
