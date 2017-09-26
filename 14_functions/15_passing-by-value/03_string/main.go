package main

import "fmt"

func main() {
	name := "Justin"
	fmt.Println(name)

	changeMe(name)
}

func changeMe(z string) {
	fmt.Println(z)
	z = "Rocky"
	fmt.Println(z)
}
