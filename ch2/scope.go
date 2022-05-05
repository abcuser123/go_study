package main

import "fmt"

func main() {
	// x := "hello!"
	// for i := 0; i < len(x); fmt.Println(i, x) {
	// 	x := x[i]
	// 	if x != '!' {
	// 		x := x + 'A' - 'a'
	// 		fmt.Printf("%c-", x)
	// 	}
	// 	i++
	// }
	//fmt.Println(i)

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c-", x)
		}
		fmt.Printf("%c\n", x)
	}

}
