package main

import "fmt"

func integer(x int) (int, bool) {

	return x / 2, x%2 == 0
}

func main() {
	var y int
	fmt.Print("Enter and Integer: ")
	fmt.Scan(&y)
	fmt.Println(integer(y))

}
