package strings

import (
	"fmt"
	"unicode/utf8"
)

// CantImmuteRunes returns and error
func CanImmuteRunes(a []rune) string {
	a[0] = 'X' //does not work
	return string(a)
}
// CantImmuteStrings returns and error
// func CantImmuteStrings(a string) string {
// 	 a[0] = 'X' //does not work
// 	 return a
// }

// RunCountDemo takes a run and byte and converts to string and shows them as runes
func RunCountDemo(ru []rune, by []byte)  {
	str := string(ru)
	s := string(by)

	fmt.Println(str, utf8.RuneCountInString(str))
	fmt.Println(s, utf8.RuneCountInString(s))
	fmt.Println(by, utf8.RuneCount(by))

}
