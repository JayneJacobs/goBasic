package interfaces

import "fmt"

// StrInterface is also an interface
type StrInterface interface {
	Myfunc() []rune
}

// StringInterface is a string type
type StringInterface string

// Myfunc is a Ingerface
func (str StringInterface) Myfunc() []rune {

	var myRune []rune
	for _, ru := range str {
		if ru == 'a' || ru == 'i' || ru == 'e' || ru == 'o' || ru == 'u' {
			myRune = append(myRune, ru)
		}
	}
	return myRune
}

// Demo uses a String type and Interface with related methods.
func Demo() {
	stringInterface := StringInterface("abcdefghi")
	var newStringInterface StrInterface
	newStringInterface = stringInterface
	myRune := newStringInterface.Myfunc()
	fmt.Printf("This is the rune ASCII value of the vowels :%v \n This is the string: %v \n", myRune, stringInterface)
}
