package interfaces

import "fmt"

// Practical is a struct interface
type Practical interface {
	GetPosition() int
}

// BandMember uses the StringInterface Type
type BandMember struct {
	Name     StringInterface
	Position int
}

// BandMemberInstrument has the BandMember type as element
type BandMemberInstrument struct {
	Member     BandMember
	Instrument string
}

// GetPosition uses the Band Member struct to return a position
func (stru BandMemberInstrument) GetPosition() int {
	return stru.Member.Position
}

// GetPosition returns the position
func (stru BandMember) GetPosition() int {
	return stru.Position
}

// BandPositionInterface ranges over all members and determines position
func BandPositionInterface(members []Practical) {
	pos := 0
	for _, i := range members {
		pos = pos + i.GetPosition()
	}
	fmt.Printf("Positions = %d\n ", pos)
}

// DemoBandInterface ranges over all members and determines position
func DemoBandInterface() {
	member1 := BandMember{"Jayne", 1}
	member2 := BandMember{"Jim", 2}
	instrument := BandMemberInstrument{member1, "Flute"}
	band := []Practical{member1, member2, instrument}
	BandPositionInterface(band)
	fmt.Printf("This is the Band: %v\n", band)
}
