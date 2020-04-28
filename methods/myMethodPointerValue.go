package methods

import "fmt"

// Placeholder is a placeholder
func Placeholder() {

	fmt.Println(" ")
}

// BandX  is a struct with a Name string and Members int
type BandX struct {
	Name    string
	Members int
}

// BandXMethodPrint Populates BandX
func (bandX BandX) BandXMethodPrint() {
	fmt.Printf(" Name: %s \n Members: %d \n", bandX.Name, bandX.Members)

}

// BandXMethodPrintRef Prints a reference to BandX
func (bandX *BandX) BandXMethodPrintRef() {
	fmt.Printf(" Name location: %d \n Members location: %d \n", &bandX.Name, &bandX.Members)
	fmt.Printf(" Name: %v \n Members: %v \n", bandX.Name, bandX.Members)
}

// BandXMethodName takes a string and and sets the bandX.name
func (bandX BandX) BandXMethodName(n string) {
	bandX.Name = n
}

// RefBandXMethodName takes a reference tothe BandX struct
func (bandX *BandX) RefBandXMethodName(s string) {
	bandX.Name = s
}

// RefBandXMethodMembers takes a reference tothe BandX struct
func (bandX *BandX) RefBandXMethodMembers(n int) {
	bandX.Members = n
}

// UseBandX uses the Method to change values in the struct fields
func UseBandX() {
	vary := BandX{
		Name:    "Asylum Wind",
		Members: 5,
	}
	vary.BandXMethodPrint()
	vary.BandMethodMembers(vary.Members + 1)
	(vary).BandXMethodName("The Asylum Wind Band")
	vary.BandXMethodPrint()
	vary.RefBandXMethodName("The Asylum Wind Band")
	vary.RefBandXMethodMembers(vary.Members + 1)
	vary.BandXMethodPrintRef()
	(&vary).BandXMethodPrintRef()
}

// BandMethodMembers takes a int and populates Members
func (bandX BandX) BandMethodMembers(num int) {
	bandX.Members = num
}

// PointerRule should be used when changes made to the receiver inside the method need to be visible to the caller.
func PointerRule() {
	fmt.Println("Pointer should be used when changes made to the receiver inside the method need to be visible to the caller.  ")
	fmt.Println("Pointer should be used it is expensive to copy a data structure. Value receivers are expensive  ")
}
