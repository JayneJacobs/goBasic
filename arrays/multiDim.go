package arrays

import "fmt"
// Threeby2 Assigns values to a copy of a 3 by 2 array
func Threeby2()  {
	a := [3][2]string {{"abc","def"},{"123","856"},{"jim","jayne"}}
	for i, v := range a {
		fmt.Printf("%d, \n %v: ", i, v)
	}
	b := a
	b[0][1]="Rich"
	b[0][0]="Rich"
	b[1][0]="Rich"
	b[1][1]="Rich"
	fmt.Println("\narray after assignment:",b)
}
// Threeby2NestedFor separates the embedded array
func Threeby2NestedFor()  {
	a := [3][2]string {{"abc","def"},{"123","856"},{"jim","jayne"}}
	for _, i := range a {
		for _, j := range i {
			fmt.Printf("\nThis is the value of the inner array elements: %s\t", j)
		}
	}

}
// NestedForArray loops through both dimentions
func NestedForArray() {
	a := [3][2]string {{"abc","def"},{"123","856"},{"jim","jayne"}}
	for _, i := range a {
		 fmt.Printf("FirstLoop = %s\t", i)
		 for _, j :=  range i {
			 fmt.Printf("In 2nd loop =%s\t", j)
		 }
		 fmt.Println("\t")
	}
}