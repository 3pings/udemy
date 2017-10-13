package main

import (
	"sort"
	"fmt"
)

func main() {

	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)
}
