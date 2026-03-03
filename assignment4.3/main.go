package main

import "fmt"

func main() {
	var arr [5]int

	fmt.Println("Enter 5 integers:")

	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr[i])
	}

	largest := arr[0]
	smallest := arr[0]

	for i := 1; i < 5; i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Smallest:", smallest)
}
