package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "Howdy"
	fmt.Println(myGreeting["Harleen"])
	myGreeting["Harleen"] = "Gidday"
	fmt.Println(myGreeting["Harleen"])
}
