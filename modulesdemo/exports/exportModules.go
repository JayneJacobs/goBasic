package exports

import "fmt"

type exported struct {
	name  string
	title string
}

func New(fi string, ll string) exported {
	ex := exported{
		fi,
		ll,
	}
	return ex
}

func (exp exported) ExportMethod() {
	fmt.Printf("This is exported feilds through Main Method %s, %s\n", exp.name, exp.title)
}
