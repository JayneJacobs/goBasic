package panicrecover

import "fmt"

// deferred functions are executed after panick
//panic / recover try / catch

//avoid panic / recover. use errors when possible

//unrecoverable errors/ programmer error . nil arguements expecting valid pointer

// func panic(interface{})  {

// }

func f1(f1 *string, l *string) {
	defer fmt.Println("\ndeferred call to a function. This can restart a program")
	if f1 == nil {
		panic("\nerror wile running f1 is nil")
	}
	if l == nil {
		panic("\nerror wile running l is nil")
	}
	fmt.Printf("going backk to caller no errors %v, %v \n", *f1, &l)
}
func recover() {
	f := "\nRestarted teh function"
	f1(&f, &f)
}

// RunPanic allows panic to print an error before exiting
func RunPanic() {
	defer recover()

	f := "Jaynee"
	// l := "Jim"
	f1(&f, nil)
	fmt.Println("\nRunPanic() ends here")
}
