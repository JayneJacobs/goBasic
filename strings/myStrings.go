package strings

import "fmt"

//Demonstrates strings as a Slice of bytes
func StringMaps()  {
	x := "This is a sentence"
	fmt.Printf(x)
	Stringbytes(x)
	hby := []byte{0x43, 0x61}
	shby := StringHex(hby)
	fmt.Printf("\nThis is the string from hex: %v \n", shby)
	aby := []byte{67, 97}
	saby := StringHex(aby)
	fmt.Printf("\nThis is the string from ascii: %v \n", saby)

}

// StringHex takes a slice of bytes and returns a string
func StringHex(x []byte) string{
	s := string(x)
	return s
}

// Stringbytes takes a string and manipulates it as a slice of bytes and shows how to print them using  verbs
func Stringbytes(a string)  {
	for i := 0; i < len(a); i++ {
		fmt.Printf("\nthis is the %vth byte: %T\t %t\t %v\t %c", i, a[i],a[i],a[i],a[i])
	}
	fmt.Printf("\nThis is using range:\t index \t ASCII \t  Value ")
	for i, val := range a {
		
		fmt.Printf("\n \t\t \t  %v\t  %v\t  %c",i, val, val)
	}
}