package main

import "fmt"

func main() {
	var arr [5]int
	sum := 0

	fmt.Println("Enter 5 integers:")

	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr[i])
		sum += arr[i]
	}

	average := float64(sum) / 5

	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)
}
