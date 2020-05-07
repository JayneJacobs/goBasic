package defering

import "fmt"

type p struct {
	f string
	l string
}

func (px p) deferedMethod() {
	fmt.Println("%s, %s", px.f, px.l)
}
func RundeferedMethod() {
	px := p{
		"Jayne",
		"Jim",
	}
	defer px.deferedMethod()
	fmt.Printf("This is in the Main Function", px)
}
