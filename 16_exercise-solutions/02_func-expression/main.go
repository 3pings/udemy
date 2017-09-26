package main

import "fmt"

func main() {
	var y int
	fmt.Print("Enter and Integer: ")
	fmt.Scan(&y)

	integer := func(x int) (int, bool) {

		return x / 2, x%2 == 0

	}
	fmt.Println(integer(y))

}
