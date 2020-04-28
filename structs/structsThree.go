package structs

import "fmt"

type Anything struct {
	a string
	b string
}

// PLaceholder is a shortcut
func PLaceholder() {
	fmt.Println(" ")
}

// BandStruct uses the Unlabled fields of a struct
func CompareStruct()  {
	Player1 := Anything{"Jayne", "Flute" }
	Player2 := Anything{"Jim", "Bass" }
	if Player1 == Player2 {
		fmt.Println("Player1 is Player 2")
		return
	}
	fmt.Println("Player1 is not Player 2\n", Player1, Player2)
	Player3 := Anything{"Jayne", "Flute" }
	Player4 := Anything{"Flute", "Jayne" }
	if Player3 == Player4 {
		fmt.Println("Player3 is Player 4")
		return
	}
	fmt.Println("Player3 is not Player 4\n", Player1, Player2)
}

// AnyStruct is a stuct of a map
type AnyStruct struct{
	Any map[int]int
}

// Comparemap shows that you cannot compare incompatible types
func Comparemap()  {
	Jayne := AnyStruct{map[int]int{0:75,}}
	Jim := AnyStruct{map[int]int{0:76,}}
   //maps are types that cannot be compared
   // if Jayne == Jim { fmt.Printline("This will never happen")}
	fmt.Println(Jayne, Jim , "Are maps and cannot be compared")
}