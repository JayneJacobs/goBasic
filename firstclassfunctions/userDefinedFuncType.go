package firstclassfunctions

import "fmt"

// MyFunc is a function type
type MyFunc func(xy int) int

// DemoUserDefinedFunctionType uses a function type
func DemoUserDefinedFunctionType() {
	var x MyFunc = func(z int) int {
		return z
	}
	g := x(111)
	fmt.Println(g)
}
