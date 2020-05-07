package mybanddesc

import "fmt"

type Myband struct {
	Name  string
	Title string
}

func (myb Myband) bandmethod() string {
	bandD := fmt.Sprintf("This is the band Name %s, Title %s", myb.Name, myb.Title)
	return bandD
}

type BandAnon struct {
	Location string
	Gig      string
	Myband
}

// BandAnonMethod exports BandAnon
func (band *BandAnon) BandAnonMethod() {
	fmt.Println(band.Location)
	fmt.Println(band.Gig)
	fmt.Println(band.bandmethod())
}
