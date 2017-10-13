package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("https://projecteuler.net/problem=13")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
