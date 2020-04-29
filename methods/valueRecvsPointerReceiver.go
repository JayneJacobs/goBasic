package methods

import "fmt"

// Value Arguements in a functions only accept value arguements
// Value Arguements in Receivers will accept a pointer or value
// Song functions and methods are different types and can have the same name
type Song struct {
	Name    string
	Minutes int
}

// Songs Prints elemnets of Song
func Songs(song Song) {
	fmt.Printf("Song1: %d, Minutes: %d \n", song.Name, song.Minutes)
}

// Songs Method Prints elemnets of Song
func (nameSong Song) Songs() {
	fmt.Printf("name: %d, minutes: %d\n", nameSong.Songs, nameSong.Minutes)
}

// SongsRef  takes a pointer to Method Prints elemnets of Song
func (nameSong *Song) SongsRef() {
	fmt.Printf("name: %d, minutes: %d\n", nameSong.Songs, nameSong.Minutes)
}

// SongsRef Prints elemnets of Song
func SongsRef(song *Song) {
	fmt.Printf("Song1: %d, Minutes: %d \n", song.Name, song.Minutes)
}

// DemoMethodvsFuction shows that you can use the same name for functions and Methods because they are different types.
func DemoMethodvsFuction() {
	song := Song{"Song sung blue", 5}
	Songs(song)
	song.Songs()
	newSong := &song

	// Songs(newSong) no can do
	Songs(*newSong)
	newSong.Songs()
	SongsRef(newSong) //can do
	newSong.SongsRef()
}
