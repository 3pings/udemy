package main

import "fmt"

//A B C exported
const (
	A = iota // 0
	B        // 1
	C        // 2
)

//D E F exported
const (
	D = iota // 0
	E        // 1
	F        // 2
)

const (
	_ = iota
	a = iota * 10
	b = iota * 20
)

//KB Exported
const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func main() {

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\n", MB)
	fmt.Printf("%d\n", MB)

}
