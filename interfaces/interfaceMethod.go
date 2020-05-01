package interfaces

import "fmt"

// BuildBand method
type BuildBand interface {
	bandMethod()
}

// BandMemberN struct
type BandMemberN struct {
	MemberName string
	MemberPos  int
}

func (bdm BandMemberN) bandMethod() {
	fmt.Printf("string is %s and int is: %d\n", bdm.MemberName, bdm.MemberPos)
}

func typeswitchBandMember(i interface{}) {
	switch b := i.(type) {
	case BandMemberN:
		b.bandMethod()
	case BandMember:
		pos := b.GetPosition()
		fmt.Println(pos)
	default:
		fmt.Printf("\nThis was an Unknown type of value:\n %d\n", b)
	}
}

// UseTypeBandMember uses the type Switch fucntion
func UseTypeBandMember() {

	myBandN := BandMemberN{
		"Asylum Wind",
		5,
	}
	typeswitchBandMember(myBandN)
	typeswitchBandMember("Asylum Wind BBand")
	myBand := BandMember{
		"Asylum Wind",
		5,
	}
	typeswitchBandMember(myBand)
}
