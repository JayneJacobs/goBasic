package interfaces

import "fmt"

// EmptyInterface (an interface that has no methods is called the empty interface
func EmptyInterface(i interface{}) {
	fmt.Printf("Type of i = %T and value = %v\n", i, i)
}

// DemoEmptyInterf shows that you can use any type
func DemoEmptyInterf() {
	Band := "Asylum Wind"
	EmptyInterface(Band)
	Bandy := "Bow & Feather"
	EmptyInterface(Bandy)
	Bands := struct {
		nameBand string
	}{nameBand: "The Asylum Wind"}
	EmptyInterface(Bands)
}

// GetUnderlyingType prints the type of the underlying string
func GetUnderlyingType(i interface{}) (string, string) {

	iString, iStype := i.(string)
	typeString := fmt.Sprintf("This is the underlying type of the string: %v and its type :%v ", iString, iStype)

	iInt, iIntType := i.(int)
	typeInt := fmt.Sprintf("This is the type of the int: %v and its type %v", iInt, iIntType)
	return typeString, typeInt
}

// DemoUnderlyingType Gets the uderlying type of the interface
func DemoUnderlyingType() {
	var Band interface{} = 64
	myBand, myInt := GetUnderlyingType(Band)
	fmt.Println(myBand, " \n", myInt)
	var Band2 interface{} = "Asylum Wind"
	myBand, myInt = GetUnderlyingType(Band2)
	fmt.Println(myBand, " \n", myInt)
}

// SimpleInterface is just that it takes any type and prints its type
func SimpleInterface(i interface{}) {
	y := i.(int)
	fmt.Println(y)
}

// ExSimpleEmptyIntf demonstrates a simple use of EmptyInterface
func ExSimpleEmptyIntf() {
	var Band interface{} = 5
	SimpleInterface(Band)

}
