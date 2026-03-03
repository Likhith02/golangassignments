package main

import "fmt"

func main() {
	var arr [5]int

	fmt.Println("Enter 5 integers:")

	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr[i])
	}

	fmt.Print("Array elements are: [")

	for i, value := range arr {
		fmt.Print(value)

		if i != len(arr)-1 {
			fmt.Print(", ")
		}
	}

	fmt.Println("]")
}
