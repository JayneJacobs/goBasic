package interfaces

import "fmt"

// TypeSwitch shows the use iwth empty interface
func TypeSwitch(band interface{}) {
	switch band.(type) {
	case string:
		fmt.Printf("The value of string is %s\n", band.(string))
	case int:
		fmt.Printf("The value of int is %d\n", band.(int))
	default:
		fmt.Printf("The value of band is not an expected type %s \n", band)
	}

}

// DemoTSwitch uses an int and a string to use the switch case
func DemoTSwitch() {
	BandName := "Asylum Wind"
	TypeSwitch(BandName)

	BandMembers := 5
	TypeSwitch(BandMembers)

	type Band struct {
		MyBandName    string
		MyBandMembers int
	}

	myBand := Band{
		"Asylum Wind",
		5,
	}
	TypeSwitch(myBand)
}
