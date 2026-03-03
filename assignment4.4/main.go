package main

import "fmt"

func main() {
	var arr [5]int

	fmt.Println("Enter 5 integers:")

	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr[i])
	}

	fmt.Println("Reversed array:")

	for i := 4; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}
