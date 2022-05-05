package main

import "fmt"

func main() {
	p := new(struct{}) //[0]int
	q := new(struct{})
	fmt.Println(p == q)

	m := new([0]int)
	n := new([0]int)
	fmt.Println(m == n)

	f()
	fmt.Println(*global)

}

var global *int

func f() {
	var x int
	x = 1
	global = &x
}
func g() {
	y := new(int)
	*y = 1
}
