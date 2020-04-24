package arrays

import "fmt"

// PrintArray Assign Elements with Index
func PrintArray(a [4]int)  {
	a[0] = 1
	a[0] = 22
	a[0] = 333
	a[0] = 4444
	fmt.Printf("elements in a %v: ", a)
}

// AssignRow TBD
func AssignRow(a [4]int)  {
	a = [4]int{ 2, 3, 4, 5}
	fmt.Printf("\nelements in a: %v", a)
}

// Findlen prints the lenth of the input array
func Findlen()  {
	a := [...]int{1,2,3,4,}
	
	fmt.Printf("\nelements in a: %v  length: %v", a, len(a))
}

//Assignment makes a copy of a variable and does not change the original value of a different name
func AssignmentCopy (){
	a := [...]string{"abc", "Jayne", "Jim"}
	b := a
	b[0] = "XXX"
	fmt.Println("This is b: ", b)
	fmt.Println("This is a: ", a)
}

// variables and elements of an array are passed by value
func PassCopy(num [4]int)  {
	num[0]=12
	fmt.Println("This is num array after assigning 12 to first element: ", num)

}

// ForArray iterates over an array by incrementing 
func ForArray(a [4]int)  {
	fmt.Println("This is teh length: ", len(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("\nThis is the %dth element %d: ", i, a[i])
	}
}

// // RangeArray iterates over an array by range function 
// func RangeArray(a [4]int) {
// 	for i, v := range a {
// 		// a[i]=a[i-1]+a[i+1]
// 		fmt.Printf("\nThis is the %dth element %d from range: ", i, v)
// 	}
// }

// RangeArray ignoring index and printing value of an array by range function 
func RangeArray(a [4]int) {
	for _, v := range a {
		// a[i]=a[i-1]+a[i+1]
		fmt.Printf("\nThis is the  element %v from range: ", v)
	}
}

