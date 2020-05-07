package structofslice

import "fmt"

type a struct {
	f string
	l string
}

func (ax a) method1() string {
	return fmt.Sprintf("%s %s \n", ax.f, ax.l)
}

type p struct {
	t string
	c string
	a
}

func (px p) method2() {
	fmt.Println(px.t)
	fmt.Println(px.c)
	fmt.Println(px.method1())
}

type w struct {
	ps []p
}

func (wx w) contents() {
	fmt.Println("contnets of wx\n")
	for _, vx := range wx.ps {
		vx.method2()
		fmt.Println(wx.ps)
	}
}

func RunsliceStruct() {
	ax := a{"abc", "def"}
	px := p{"123", "456", ax}
	px1 := p{"Jayne", "Jim", ax}
	wx := w{
		[]p{px, px1},
	}
	wx.contents()
}
