package interfaces

import "fmt"

// StringInterface is also an interface
type StrInterface interface {
	myfunc() []rune
}

// StringInterface is a string type
type StringInterface string

func (str StrInterface) Myfunc() []rune {

	var myRune []rune
	for _, ru := range str {
		if ru == 'a' || ru == 'i' || ru == 'e' || ru == 'o' || ru == 'u' {
			myRune = append(myRune, ru)
		}
	}
	return myRune
}

func InterfacesDemo() {
	stringInterface := StringInterface("abcdefghi")
	var newStringInterface StrInterface
	myRune := newStringInterface.Myfunc()
	fmt.Printf("This is the rune :%v \n This is the string: %v \n", a ...interface{})
}
