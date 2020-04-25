package slices

import "fmt"

func MultDimSlice(a [][]string) {
	 for _, k := range a {
		for _, j := range k {
			fmt.Printf("%s ", j)
			fmt.Sprintf("%s ", k)

		}
		fmt.Println("\n")
	}
	
}