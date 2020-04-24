package panics

import (
	"fmt"
	"time"
)

// PanicEx print statemens and show Panic also User Input Example
func PanicEx()  {
	var a int
	fmt.Println("enter1 for amatuer 2 for expert: ")
	fmt.Scanln(&a)
	switch a {
	case 1:
		fmt.Println("Amateur")
	case 2:
		fmt.Println("Expert")
	default:
		panic(fmt.Sprintf("Entered something incorrectly %d", a))
	}
   fmt.Println("Back in main This wont print if Panic runs")
  
}

// PanicDefer demonstrates a panic in an if statement// pointer to a string
func PanicDefer (x *string, y *string) {
	defer RecoverPanic()
	defer fmt.Println("\nThis will print if Panic runs, it is defered")
	
	if x == nil {
		panic("\nruntime err: x is nil")
	}
	if y ==nil {
		panic("\nruntime err: y is nil")
	}
	fmt.Printf("\nThis is x: %s \n This is y %s", *x,*y)
}

// RecoverPanic allows to panic from a function and return to another function
func RecoverPanic()  {
	if r := recover(); r != nil {
		fmt.Println("\nrecovered from: ", r)
	}
}
// PanicFirst calls recoverPanic and runs PanicSecond as a go routine since panic happens in PanicSecond it will not recover
func PanicFirst() {
	defer RecoverPanic()
	fmt.Println("in PanicFirst")
	// go PanicSecond() 
	//removing go from PanicSecond allows recovery
	PanicSecond()
    time.Sleep(1 * time.Second)

}

func PanicSecond(){
	fmt.Println("in PanicSecond")
	panic("Panicking in Panic second")
}