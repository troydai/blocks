package main

import "fmt"

func main() {
	closure()
}

func closure() {
	word := "sample"

	defer func(input string) {
		fmt.Print("case 1: ")
		fmt.Println(input)
	}(word)
	defer func() {
		fmt.Print("case 2: ")
		fmt.Println(word)
	}()

	word = "sample-changed"
	defer func(input string) {
		fmt.Print("case 3: ")
		fmt.Println(input)
	}(word)
}
