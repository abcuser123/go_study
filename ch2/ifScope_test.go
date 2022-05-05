package ch2

import (
	"fmt"
	"testing"
)

func Test_if(t *testing.T) {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := 2; x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
}

func f() int {
	return 1
}
