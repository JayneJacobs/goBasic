packge pointref
imports(
   "fmt"
)
//A pointer stores the memory location of a variable.  

// PointerArrayDemo of passing by pointer changing original value in another function
func PointerArrayDemo() {
	x := [3]int{89, 92, 96}
   fmt.Printf("\nThis is the original Value of x: %v \n" ,x)

   y := PointArray(&x)

   fmt.Printf("\nThis is the  Value of x after passing the address as &x to a function: %v \n", x)
   fmt.Printf("\nThis is the original Value of y returyne by the function: %v \n", y)
   y = PointArray(&x)
   fmt.Printf("\nThis is the  Value of x after passing the address as &x to a function again: %v \n", x)
   fmt.Printf("\nThis is the original Value of y returyne by the function again:  %v\n", y)
   z := *(&x)
   fmt.Printf("This is a dereferenced address: %v \n", z)
   z = x
   fmt.Printf("This is a  address: %v  Type %T\n", z, z)
}


// PointArray demonstrates passing a  an address of an array and using more or less *&var 
func PointArray(a3 *[3]int) [3]int{
(*a3)[2]=(*a3)[2]+125
a3[2]=(*a3)[2]+125
return *a3
}

// PointerDemo shows use of &Address of var *Reference to value at address
// func Pointer(a *int, b int) *int{

// fmt.Printf("\nThis is the address of reference to a: %v \n", a)
// fmt.Printf("\nThis is the address of b: %v \n", &b)
// b = b + 100 //has no effect on the orginal value of b in main
// fmt.Printf("\nThis is the Type of a %T: ", a)
// fmt.Printf("\nThis is the dereference  value stored at the address of *a %d: ", *a )
// *a = *a +100
// return a

}
