package panicruntime

import (
	"fmt"
	"runtime/debug"
)

//panics caused by runtime errors such as nil pointer to struct

func x() {

	defer recoverfunc()
	nx := []int{8, 7, 3}
	fmt.Println(nx[3])
	fmt.Println("successfully returned to main")
}

func recoverfunc() {
	if rx := recover(); rx != nil {
		fmt.Println("recovered from ", rx)
		debug.PrintStack()
	}
}

// CallMain shows a runtimepanic
func CallMain() {

	x()
	fmt.Println("end of Main")
}
