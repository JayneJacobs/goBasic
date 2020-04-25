package slices

import "fmt"

//A reference to an array

// AssigntoSlice takes an array and assigns it to a slice
func AssigntoSlice(a [5]int)  {
	b := a[1:4] //a of start to a n-1
	fmt.Printf("\nThis is 1:4 of a as a slice %v", b)
}


func ModifySlice (a [6]int){
	b := a[1:4] //a of start to a n-1
	fmt.Printf("\nThis is 1:4 of a as a slice b %v", b)
	fmt.Printf("\nThis is 1:4 of a array a %v", a)
	
	for i := range b {
          b[i]++
	}
	
	fmt.Printf("\n\n\nThis is 1:4 of a as a slice b after for loop b[i]++ %v", b)
	fmt.Printf("\nThis is 1:4 of a as a slice a  after for loop %v", a)
	b1 := a[1:4]
	for i := range b1 {
		b1[i]++
  }
	fmt.Printf("\n\n\nThis is 1:4 of a as a slice b after for loop b1[i]++ %v", b)
	fmt.Printf("\nThis is 1:4 of a as a slice b1  after for loop %v", b1)
	fmt.Printf("\nThis is 1:4 of a as a slice a  after Second for loop %v", a)

}


func ModifyCapSlice (a [6]int){
	b := a[1:4] //a of start to a n-1
	fmt.Printf("\nThis is 1:4 of a as a slice b %v", b)
	fmt.Printf("\nThis is 1:4 of a array a %v", a)
	
	for i := range b {
          b[i]++
	}

	fmt.Println("\nSlice operations share the values of the underlying arrays")
	fmt.Println("\nCapacity = num elements in underlying array starting from the index \n from which the slice is created")

	fmt.Println("\nLen ",len(b))
	fmt.Println("\nCapacity ",cap(b))

	b = b[:cap(b)]
	fmt.Println("\nLen after settting b to cap(b) ",len(b))
	fmt.Println("\nCapacity after settting b to cap(b) ",cap(b))
	fmt.Printf("\n\nThis is slice b  %v   after settting b to cap(b)", b)
	fmt.Printf("\nThis is a :%v     aftersettting b to cap(b)", a)
}

func ReSlice (a [6]int){
	b := a[1:4] //a of start to a n-1
	fmt.Printf("\nThis is 1:4 of a as a slice b %v", b)
	fmt.Printf("\nThis is 1:4 of a array a %v", a)
	
	
	fmt.Println("\nLen ",len(b))
	fmt.Println("\nCapacity ",cap(b))

	b = b[:cap(b)]
	fmt.Println("\nLen after settting b to cap(b) ",len(b))
	fmt.Println("\nCapacity after settting b to cap(b) ",cap(b))
	fmt.Printf("\n\nThis is slice b  %v   after settting b to cap(b)", b)
	fmt.Printf("\nThis is a :%v     aftersettting b to cap(b)", a)
}

