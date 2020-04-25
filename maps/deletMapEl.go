package maps

import "fmt"

// DeleteKey deletes teh key that is passed and returns the resulting map
func DeleteKey(x map[string]string, key string ) (map[string]string,bool, string) {
	delete(x, key)
	val, ok := x[key]
	if ok {
		return x,ok, val
	}
	for key, val := range x {
		delete(x, key)
		fmt.Printf("\nThis Key Exists: %v, as %v \n", key, val)
		delete(x, key)
	}
	if ok {
		delete(x, key)
		return x,ok, val
	}
	return x, ok, "not presnet"
}
// CkPStringforKey determines if a key exists
func CkPStringforKey(x map[string]string, key string ) (bool, string) {
	val, ok := x[key]
	if ok {
		return ok, val
	}
	for key, val := range x {
		fmt.Printf("\nThis Key Exists: %v, as %v \n", key, val)
		delete(x, key)
	}
	if ok {
		return ok, val
	}
	return ok, "not presnet"
}
