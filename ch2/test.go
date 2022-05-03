package main

import (
	"fmt"
)

func main() {
	//关键字
	//interface select defer chan goto const fallthrough range

	//预声明
	//类型
	//true false iota nil  常量
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//float32 float64  complex128 complex64
	//bool byte rune string error
	//函数
	//make len cap new append copy close delete
	//complex real imag
	//panic recover

	fmt.Print("G")
	var temp string
	fmt.Print(temp)
	fmt.Println("O")

	i, j := 0, 1
	i, j = j, i
	fmt.Print(i, j)
	i, k := 2, 3
	fmt.Println(i, k)

	// f,err := os.Open(name)
	// if err!=nil{
	// 	return err
	// }
	// f.Close

	var p = f()
	fmt.Println(p)
	fmt.Println(f() == f())

	v1 := 1
	incr(&v1)
	fmt.Println(incr(&v1))
	//无区别
	v := 1
	incr(&v)
	fmt.Println(incr(&v))

}
func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
