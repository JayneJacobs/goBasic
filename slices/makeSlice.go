package slices

import "fmt"

// MakeSlices demonstration
func MakeSlices()  {
	k := make([]int, 4, 5) //len, cap
	fmt.Printf("This makes an underlying array and a slice reference to the array indexes: %v",k)

}


// AppendSlices with a slice
func AppendSlices(a []string)  {
	l := len(a)
	c := cap(a)
	  fmt.Printf("This is the initial slice a:\n  len :  %v \n cap: %v ",l,c)
	  a = append(a, "Rich")
	  l = len(a)
	c = cap(a)
	  fmt.Printf("\nThis is the appended value of slice a: %v \n len: %v \n cap: %v\n",a,l,c)
       fmt.Println("When the slice is appended a new array is created the elements are copie to the original slice reference it is 2x the old slice")
	}

	// NilSlices appends the initial nil slice with a slice 
	func NilSlices(a []string)  {
	var n[] string
	fmt.Printf("The nil slice: %v", n)
	l := len(n)
	c := cap(n)
	fmt.Printf("This is the initial slice n:\n  len :  %v \n cap: %v ",l,c)
	for _, v := range a {
		a = append(n, v)
	}
	l = len(a)
	
	  c = cap(a)
	  fmt.Printf("\nThis is the appended value of slice n : %v \n len: %v \n cap: %v\n",a,l,c)
       fmt.Println("When the slice is appended a new array is created the elements are copie to the original slice reference it is 2x the old slice")
}
// AppendSlicedot appends a slice with values from a sliceusing ...
func AppendSlicedot(a []string)  {
	b := []string{"Rich","Scott", "Elvis"}
	c := append(a, b...)
	fmt.Printf("\nThis is struct c the appended a with elemnets of b using a, b...: %v", c)
	fmt.Printf("\nThis is the struct a appended a with elemnets of b using a, b...: %v", a)
	a = append(a, b...)
	fmt.Printf("\nThis is the struct a appended a with elemnets of b using a, b...: %v", a)
	l := len(a)

	cp := cap(a)
	fmt.Printf("\nThis is the appended value of slice n : %v \n len: %v \n cap: %v\n",a,l,cp)
}

// Funk takes a slice appends it and returns an element of the slice
func Funk(n []int) []int {
	for i := range n {
		n[i]= -2
	}
	return n
}


