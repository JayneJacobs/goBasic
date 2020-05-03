package interfaces

import "fmt"

// NilIntf uses the Discover method
type NilIntf interface {
	NewDiscover()
}

// DiscIntf uses the Discover method
type DiscIntf interface {
	Discover()
}

// Myls is an interface for the EMember struct that returns a int
type Myls interface {
	ReturnIntMethod() int
}

// EmbedNilandMyls uses both interfaces for Emember Type
type EmbedNilandMyls interface {
	DiscIntf
	Myls
}

// EMember is a band struct
type EMember struct {
	FirstName  string
	LastName   string
	Instrument string
	Position   int
	Age        int
	Songs      int
}

// Discover prints a string and an int from the EMember struct
func (emember EMember) Discover() {
	fmt.Printf("Firstname\tLastname\tPosition\n%v\t\t%v\t\t%d\n", emember.FirstName, emember.LastName, emember.Position)
}

// ReturnIntMethod uses Age from EMember
func (emember EMember) ReturnIntMethod() int {
	return emember.Age
}

// DemoMultipleInterfaces uses two interfaces on the same type
func DemoMultipleInterfaces() {
	embr := EMember{
		"Jayne",
		"Jacobs",
		"Flute",
		1,
		62,
		100,
	}
	var embr2 DiscIntf = embr
	embr2.Discover()

	var myls Myls = embr
	fmt.Printf("This is the Second Interface returns Age: %d \n", myls.ReturnIntMethod())

	var embedintf EmbedNilandMyls = embr
	embedintf.Discover()
	fmt.Printf("This is Printed Emember Age using Embeded interface: %d \n", embedintf.ReturnIntMethod())

	//FOR THE NIL INTERFACE
	var newNil NilIntf
	if newNil == nil {
		fmt.Printf("This is the Nil interace:\nType\tValue\n%T\t%v\n", newNil, newNil)
		fmt.Println("The Nil Interface has an underlying value not a concrete type ")
	}
	// newNil.NewDiscover() //Nocan do this will panic if it is a nil Type
}
