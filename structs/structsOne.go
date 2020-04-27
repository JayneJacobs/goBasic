package structs

import "fmt"

// RefStructDemon
func RefStructDemo() {
	x := &Band{"Jayne", "Flute", 1} //Needs all three values
	fmt.Printf("Name: %v\n", x.Name)
	fmt.Printf("Name: address: %p, value: %v", &(x).Name, (*x).Name)

}

// StructDemo demonstrates how to access and assign values to a struct
func StructDemo() {
	x, y := DemoStruct()
	fmt.Println(x)
	fmt.Println(x.Position)
	fmt.Println(y.Position, y.Name)
	z, q := RefStruct()
	fmt.Println("This is the initial values of z", z.Name, z.Position)
	z.Name = "Jayne"
	z.Instrument = "Keyboard"
	fmt.Println("These are the values of z after changing them: ", z.Name, z.Instrument)
	fmt.Println(q) //Empty
	fmt.Println(q.Name, q.Instrument, q.Position)
	q.Instrument = "Guitar"
	q.Name = "Rich"
	fmt.Printf("This is the final struct after Assigning values to  q:\n\t %v \n Its type is: \t%T \n and address: %p \t\n", *q, *q, q)
	fmt.Printf("This is z: \t %v \n and its address: \t %p\n", z, z)
}

// Band is a struct with strings
type Band struct {
	Name       string
	Instrument string
	Position   int
}

// RefStruct two ways to assign values to structs
func RefStruct() (*struct {
	Instrument string
	Name       string
	Position   int
}, *Band) {
	z := &struct {
		Instrument, Name string
		Position         int
	}{
		Name:       "Jayne",
		Instrument: "guitar",
		Position:   1,
	}
	var q Band //No values assigned

	fmt.Printf("This is the memory location of q1:  %p \n", &q)
	return z, &q

}

// DemoStruct two ways to assign values to structs
func DemoStruct() (Band, Band) {
	x := Band{
		Name:       "Rich",
		Instrument: "Guitar",
		Position:   3,
	}
	y := Band{
		"Jayne",
		"Fluet",
		1,
	}
	return x, y

}
