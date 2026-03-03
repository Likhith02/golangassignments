package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)
	result := add(a, b)
	fmt.Println("Sum is : ", result)

}
func add(x int, y int) int {
	return x + y
}
