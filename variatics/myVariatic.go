package variatics

import (
	"fmt"
	"reflect"
)

//Ve is a VariaticFunction
func Ve(s ...string) {
	fmt.Printf("This is the 1st string: %v %d", s[1], &s[1]) 
	fmt.Printf("\nThis is the 2nd string: %v \n This is the array %v: ", s[2], s) 
   }
   // Vn is a variatic function 
   func Vn(s ...string) {
	   fmt.Printf("\n This is the empty string: %v %s", &s, &s)
   }
   
   // VindNstring int and variatic args
   func VindNstring(x int, s ...string)  {
	   fmt.Printf("\nThis is s :%v ", s)
   }
   
   // VariInterf prints the types in the empty interface
   func VariInterf(i ...interface{})  {
	   fmt.Printf("\nThis is the variatic Interface: %v", i)
	   for y, v := range i {
		   fmt.Printf("\nThis is the use of empty interfac to use varied types: %v %v, %v", y, reflect.ValueOf(v).Kind(), v)
		   // fmt.Printf("\nThis is the use of empty interfac to use varied types: %v", y)
	   }
	   
   }