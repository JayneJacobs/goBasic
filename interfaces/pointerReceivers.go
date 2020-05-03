package interfaces

import "fmt"

// BandIntf has a bandMethod
type BandIntf interface {
	BandMethod()
}

// BandMemNames is  a list of Member Names
type BandMemNames struct {
	Name1 string
	Name2 string
	Name3 string
	Name4 string
	Name5 string
}

// BandNamePos combbin
type BandNamePos struct {
	Members BandMemNames
	Count   int
}

// BandMethod Builds the struct
func (bandNamePos *BandNamePos) BandMethod() {

	fmt.Printf("This is the Band: %v\n", bandNamePos)

}

// NoPointerMethod Demonstrates a nonPointer receiver
func (bandNamePos BandNamePos) NoPointerMethod() {
	fmt.Printf("This is the band wih no pointer: %v \n", bandNamePos)
}

// DemoPointerReceiver builds BandNamePos Struct
func DemoPointerReceiver() {
	var myBandPos BandMemNames
	myBandPos.Name2 = "Jim"
	myBandPos.Name1 = "Jayne"
	var bandNamePos BandNamePos
	bandNamePos.Count = 5
	bandNamePos.Members = myBandPos
	bandNamePos.BandMethod()

	myBandPos2 := &BandMemNames{"Jayne", "Jim", "Rich", "Elvis", "Scott"}
	var bandNamePos2 BandNamePos
	bandNamePos2.Members = *myBandPos2
	bandNamePos2.Count = 5
	bandNamePos2.BandMethod()
	// BandMethod()
}
