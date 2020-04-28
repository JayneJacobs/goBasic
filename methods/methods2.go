package methods

import "fmt"

// BandMember has a description of a Person with Name, Instrument and Position
type BandMember struct {
	Name       string
	Instrument string
	Position   int
}

// PrintMember Prints Name and Position
func (bandM *BandMember) PrintMember() {
	fmt.Printf("Name: %s\n Position: %s\n", bandM.Name, bandM.Instrument)
}

// Members a list of members in a slice
type Members struct {
	Bandmembers []BandMember
}

// Band is the list of members and band name
type Band struct {
	BandName    string
	MemberNum   int
	MembersList Members
}

// AssignBandmembers assigns and prints the band result
func AssignBandmembers() {
	bandM := BandMember{
		"Jayne",
		"Flute",
		1,
	}
	bandM2 := BandMember{
		"Jim",
		"Bass",
		2,
	}
	members := Members{
		[]BandMember{bandM, bandM2},
	}
	band := Band{
		BandName:    "Asylum Wind",
		MemberNum:   5,
		MembersList: members,
	}
	(&band).PrintBand()
	(&bandM).PrintMember()
}

// PrintBand brints the band a Name and members
func (band *Band) PrintBand() {
	fmt.Printf("This is the band name: %v\n This is the Member List: %v \n", &band.BandName, &band.MembersList)
}
