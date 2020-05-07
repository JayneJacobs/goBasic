package errorsl

import (
	"fmt"
	"path/filepath"
)

// DemoFilePath compares err to BadPaternError
func DemoFilePath() {
	f, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
