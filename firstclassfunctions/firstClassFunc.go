package firstclassfunctions

import "fmt"

// firstclass functions can be named as a variable and returned from a function.
//Anonymious functions

// DemoFirstClassFunc calls an anonymous function
func DemoFirstClassFunc() {
	x := func() {
		fmt.Println("hi", nil, "There")
	}
	x()
	fmt.Printf("This is the type  of x:  %T", x)
	y := " \nCall without variables."
	DemoCallAnonFuncsinVariable(y)
}

// DemoCallAnonFuncsinVariable pass arguments to anonymous functions
func DemoCallAnonFuncsinVariable(x string) {

	func(x string) {
		fmt.Printf("\nHi %v\n", x)
	}(x)

}

//User defined function types
