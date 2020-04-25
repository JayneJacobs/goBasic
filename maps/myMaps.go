package maps

import "fmt"

// CheckMapKey takes a map and key and returns a state for the key
func CheckMapKey(x map[string]int, key string)( bool, string, int ){

	val, ok := x[key]
	fmt.Println(val)
	if ok {
		return ok, key, val
	}
	return false, key, val
}

// AssignKey returns the result of x[key]
func AssignKey(x map[string]int, key string) (int, string ){
	 a := x[key]
	 b := fmt.Sprint(x)
	 return a, b
	
}



  func MakeAssignMap() (map[string]int, string) {
   
	x := map[string]int {
		"Jayne":1,
		"Jim":2,
	}
	for i, j  := range x {
		a := fmt.Sprint(x[i])
		fmt.Println(j)
		return x, a
	}
	return x, ""

}

// UseMaps returns a string result of a map manipulation
func UseMaps(x map[string]int) (string, string){
	b := fmt.Sprintln(x)
	if x == nil {
		b := fmt.Sprintln("x is nil")
		return b, ""
	}
	x["\tJayne"] = 1
	x["\n\tJim"] = 2
	x["\n\tRich"] = 3
	a := fmt.Sprintln(x)
	return b, a
}
