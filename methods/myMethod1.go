package methods

import "fmt"

// Placeholder is a placeholder
func Placeholder() {

	fmt.Println(" ")
}

type BandX struct {
	Name    string
	Members int
}

// BandXMethod1 Populates BandX
func (bandX BandX) BandXMethod1() {
	fmt.Printf(" Name: %s \n Members: %d \n", bandX.Name, bandX.Members)

}

// BandXMethod2 takes a string and and sets the bandX.name
func (bandx BandX) BandXMethod2(n string) {
	bandx.Name = n
}

// UseBandX uses the Method to change values in the struct fields
func UseBandX() {
	vary := BandX{
		Name:    "Asylum Wind",
		Members: 5,
	}
	vary.BandXMethod1()
}
