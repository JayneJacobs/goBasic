package channels

import (
	"fmt"
	"sync"
	"time"
)

//

func waitGroups(y int, ab *sync.WaitGroup) {
	fmt.Println("go routine", y, "Begining")
	time.Sleep(2 * time.Second)
	fmt.Printf("Go routine %d ended \n", y)
	ab.Done()
}

// RunWaitGroup
func RunWaitGroup() {
	v := 3
	var ab sync.WaitGroup
	for t := 0; t < v; t++ {
		ab.Add(1)
		go waitGroups(t, &ab) // pass address so it is not a copy
	}
	ab.Wait() //Blocks untill counter becomes 0
	fmt.Println("all go routines complete")
}
