package structs

import (
	"fmt"
	"time"
)

// DemoPromoted uses Promoted fields of an anonymous struct
func DemoPromoted() {
	var y Instruments
	y.Person1.Instrument = "Flute"
	y.Person1.Name = "Jayne"
	fmt.Printf("This is the Instrument description: %v", y)

	var z Bands
	First := z.Band1.Person1
	z.Gig = "Philadelphia"
	fmt.Printf("This is Band 1 : \n\t%v Gig: \n\t %v", First, z.Gig)

	var p Anonymous
	p.int = 1
	p.string = "Jayne"
	fmt.Printf("This is the promoted fields: \nAnonymous: %v", p)
}

//Anonymous fields struct
type Anonymous struct {
	int
	string
}

// BandMembers Name string and Position int
type BandMembers struct {
	Name       string
	Position   int
	Instrument string
}

// Instruments is a struct describing a Band
type Instruments struct {
	Person1  BandMembers
	Person2  BandMembers
	Person3  BandMembers
	Location string
	GigDate  time.Weekday
}

// Bands describes a group of bands at a gig location
type Bands struct {
	Band1 Instruments
	Gig   string
}
