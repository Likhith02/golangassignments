package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter a number: ")
	fmt.Scanln(&num)
	result := 1
	for i := 1; i <= num; i++ {
		result = result * i
	}
	fmt.Println("factorial is:", result)
}
