package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "Howdy"
	fmt.Println(myGreeting)
	delete(myGreeting, "Harleen")
	fmt.Println(myGreeting)

	if val, exists := myGreeting["Harleen"]; exists {
		fmt.Println("That value Exists!")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value does not exists")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)

	}
}
