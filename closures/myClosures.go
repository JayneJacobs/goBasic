package closures

import "fmt"

//Special case of anonymousfuction

// DemoClosures shows using surrounding var where the func is run
func DemoClosures() {
	x := 10
	func() {
		fmt.Println("x= ", x)
	}()
}

func x() func(string) string{
	xy := "Hello"
	ct := func (xb string) string {
		xy = xy + " + xb"
		return xy
	}
	return ct
}

// DemoClosure in a function that returns an anonymousfuction
func DemoClosureinCalledAnonFunct()  {
	ax := x()
	ay := x()
	fmt.Println(ax("this is what was printed in ax"))
	fmt.Println(ay("eachone ay"))
	fmt.Println(ay("hi ay"))
}