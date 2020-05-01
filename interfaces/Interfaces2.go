package interfaces

import "fmt"

// Band is an interface needing the Band function
type Band interface {
	band()
}

// Position is a float64
type Position float64

func (p Position) band() {
	fmt.Printf("This is my band: %v\n", p)
}

func typesInterface(b Band) {
	fmt.Printf("This is the type %T of Interface band: %v \n", b, b)
}

// ImplementInterface used band function
func ImplementInterface() {
	var b Band
	myband := Position(10.2) //Position is the concrete type of Band
	b = myband
	typesInterface(b)

}
