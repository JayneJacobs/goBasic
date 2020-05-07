package errorsl

import (
	"fmt"
	"net"
	"os"
)

// ShowFileOpenErrors opening File
func ShowFileOpenErrors() {
	x, err1 := os.Open("errors/myfiley.txt")
	if err1 != nil {
		fmt.Println("Error Opening File: ", err1)
		return
	}
	fmt.Println(x.Name(), "successfully opened")
}

func AnotherFileOpenError() {
	x, err1 := os.Open("./errors/myfiley.txt")
	if err1, ok := err1.(*os.PathError); ok {
		fmt.Printf("File at path", err1.Path, "failed to open", x)
	}
}

func ThirdErrorDNSLookup() {
	x, err1 := net.LookupHost("www.go0000oogle.com")

	if err1, ok := err1.(*net.DNSError); ok {
		if err1.Timeout() {
			fmt.Println("operation has timed out")
		} else if err1.Temporary() {
			fmt.Println("temp error")
		} else {
			fmt.Println("generic err", err1)
		}
		return
	}
	fmt.Println(x)
}

type error interface {
	Error() string
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " :" + e.Err.Error()
}
