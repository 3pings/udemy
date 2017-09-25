package main

import "fmt"

func main() {
	var lgNum int
	var smNum int
	fmt.Print("Enter a large number: ")
	fmt.Scan(&lgNum)
	fmt.Print("Enter a small number: ")
	fmt.Scan(&smNum)
	fmt.Println(lgNum, "/", smNum, "=", lgNum/smNum)
}
