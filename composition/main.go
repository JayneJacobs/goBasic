package main

import "goBasic/composition/mybanddesc"

func main() {
	band := mybanddesc.Myband{
		"Jayne",
		"Floutist",
	}
	thisBand := mybanddesc.BandAnon{
		"Philadelphia",
		"Frongs",
		band,
	}
	thisBand.BandAnonMethod()
}
