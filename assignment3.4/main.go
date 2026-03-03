package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	result := max(a, b)

	fmt.Println("Maximum number is:", result)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
