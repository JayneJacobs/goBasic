package main

import "fmt"

func main() {
	// variatics.Ve("abc", "jayne", "jim")
	// variatics.Vn()
	// variatics.VindNstring(1, "hi")
	// variatics.VariInterf(1, "abbc", true, 12.12, []string{"123","xxx", "Jayne"})
// PanicEx()
	// n := 1
	// prime := prime.IsPrime(n)
	// fmt.Println(prime)
	// prime2 := prime.PrimeCk(n)
	// fmt.Println(prime2)
//Panic Part1
	// defer fmt.Println("This is in main, it is defered")
	// // x:= "abc"
	// // PanicDefer(&x,nil)
	// panics.PanicFirst()
	// fmt.Println("Back in Main funciton after panic which does not happen if panic runs")
	// var a [4]int
//Slices
	// arrays.PrintArray(a)
	// arrays.AssignRow(a)
	// arrays.Findlen()
	// arrays.AssignmentCopy()
	// num := [...]int{01,856,275,3638,6908,6926}
	// arrays.PassCopy(num)
	// fmt.Println("This is num array after passing it by value(copy), still unchanged: ", num)
	// arrays.ForArray(num)
	// arrays.RangeArray(num)
	// arrays.Threeby2()
	// arrays.Threeby2NestedFor()
	// arrays.NestedForArray()
// Slices
	// slices.AssigntoSlice(num)
	// slices.ModifSlice(num)
	// slices.ModifyCapSlice(num)
	// slices.ReSlice(num)
	// fmt.Printf("\nThis is num in main  after Second for loop %v", num)
	// slices.MakeSlices()
	// slc := []string{"Jayne","Jim"}
	// slices.AppendSlices(slc)
	// slices.NilSlices(slc)
	// a := []int{1,2,3}
	// fmt.Println(a)
	// a = slices.Funk(a)
//MultiDimentional slices need commas after each line
//     a := [][]string {
// 		{"Jayne","Jim"},
// 		{"Scott", "Rich", "Elvis"},
// 	}
//    slices.MultDimSlice(a)
//maps 
	// x := make(map[string]int) 
// var x map[string]int
// UseMaps returns a string
// 	  mys, myy := maps.UseMaps(x)
// 	  fmt.Printf("\nBefore assign: \n %v ",mys)

// 	  mys, myy  = maps.UseMaps(x)
// 	  fmt.Printf("\nAfter assign: \n %v   ",myy) // need the make funciton for x
// // Assignement
// 	x, j := maps.MakeAssignMap()
// 	fmt.Println(x["Jayne"])
// 	x["Jayne"] = 5
// 	x["Jim"] = 2
// 	x["Rich"] = 3
// 	fmt.Println( j, x["Rich"], x["Jayne"])
// //Assign Key for Maps
// 	key := "New"
// 	a, b := maps.AssignKey(x, key)
// 	fmt.Println( a, b)
// 	exist, key, val := maps.CheckMapKey(x,key)
// fmt.Printf("/This key: %v exists: %v: it is %v ", key, exist, val)
// x := make( map[string]string)
//     key1 := "Jayne"
// 	x[key1] = "Flute"
// 	x["Jim"] = "Bass"
// 	x["Rich"] = "Guitar"
//     key2 := "new"
// 	ok, val :=  maps.CkPStringforKey(x, key1)
// 	fmt.Printf("\nThis value %v, for thi key: %v: %v\n",val, key1, ok)
// 	ok, val =  maps.CkPStringforKey(x, key2)
// 	fmt.Printf("\nThis value %v, for this key: %v: is it there? %v \n",val, key2, ok)
// x := make( map[string]string)
//     key1 := "Jayne"
// 	x[key1] = "Flute"
// 	x["Jim"] = "Bass"
// 	x["Rich"] = "Guitar"
//     key2 := "new"
// 	y,ok, val :=  maps.DeleteKey(x, key1)
// 	fmt.Printf("\nThis value %v, for thi key: %v: %v\n",val, key1, ok)
// 	y,ok, val =  maps.DeleteKey(x, key2)
// 	fmt.Printf("\nThis value %v, for this key: %v: is it there? %v \n",val, key2, ok)
// 	fmt.Printf("This is the new map: %v",y)
//Maps with strings
	// strings.StringMaps()
	// ru := []rune{0x0053, 0x0065, 0x00f1}
	// x := []byte{0x54, 0x61}
	// strings.RunCountDemo(ru, x)
// Are strings immutable
// x := "hihello"
// 	// err := strings.CantImmuteStrings(x)
// 	// fmt.Println(err)
// 	err := strings.CanImmuteRunes([]rune(x))
// 	fmt.Println(err)
//Pointers
// b := 100
// a := &b
// an := PointerDemo(a, b)//A is the address of b 
// fmt.Printf("\nThis is now the address of a: %v, \t new value of b: %v ", an, b)
// an = PointerDemo(&b, b) //&b is also the address of b. 
// fmt.Printf("\nThis is now the address of a: %v, \t new value of b: %v ", an, b)
} 

PointArray(){
	
}

// PointerDemo shows use of &Address of var *Reference to value at address
func PointerDemo(a *int, b int) *int{

fmt.Printf("\nThis is the address of reference to a %v: ", a)
fmt.Printf("\nThis is the address of b %v: ", &b)
b = b + 100 //has no effect on the orginal value of b in main
fmt.Printf("\nThis is the Type of a %T: ", a)
fmt.Printf("\nThis is the dereference  value stored at the address of *a %d: ", *a )
*a = *a +100
return a

}
