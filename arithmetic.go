package main

import "fmt"

func main() {
	fmt.Println("Enter digit 1: ")
	var digit1 int
	fmt.Scanln(&digit1)
	
	fmt.Println("Enter digit 2: ")
	var digit2 int
	fmt.Scanln(&digit2)

	fmt.Println("Enter operator: ")
	var operator string
	fmt.Scanln(&operator)

	var answer int = 0
	if operator == "+" {
		answer = digit1 + digit2
	} else if operator == "*" {
		answer = digit1 * digit2
	} else if operator == "-" {
		answer = digit1 - digit2
	} else if operator == "/" {
		answer = digit1 / digit2
	} else {
		fmt.Println(fmt.Errorf("math: invalid operator"))
	}
	fmt.Println("The answer is:", answer)
	
}