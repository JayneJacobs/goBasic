package areas

import (
	"fmt"
	"goBasic/errorsl"
	"math"
)
//Using an ErrorStruct rules must contain string named error
type E struct {
	err string
	r float64
}

//Error type must have a method Error that has a *Error receiver and returns a string
func (ex *E) Error() string {
	return fmt.Sprintf("%f %s", ex.err)
}

func careaErrType(r float64) (float64, error) {
	if r < 0 {
		return 0, &E{"calc failed", r}
	}
	return math.Pi * r * r, nil
}


func carea(r float64) (float64, error) {
	if r < 0 {
		return 0, errorsl.New("calc failed")
	}
	return math.Pi * r * r, nil
}

func careaErrorf(r float64) (float64, error) {
	if r < 0 {
		return 0, fmt.Errorf("\ncalc failed%0.2f", r)
	}
	return math.Pi * r * r, nil
}
// Carea will return an error due to negative radius
func Carea()  {
	r := -20.0
	a, err := carea(r)
	if err != nil {
		fmt.Println(err)
		
	}
fmt.Printf("%f", a)

b, err1 := careaErrorf(r)
if err1 != nil {
	fmt.Println("\nFrom Errorf", err1)
	return
}
fmt.Printf("%f", b)
c, err2 := careaErrorf(r)

if err2 != nil {
	fmt.Println("\nFrom Errorf", err2)
	return
}
fmt.Printf("\n%f in error message\n", c)
}