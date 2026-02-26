package main

import "fmt"

func largest() {
	var a, b, c int
	fmt.Print("Enter the 1st number : ")
	fmt.Scanln(&a)

	fmt.Print("ENter the 2nd number: ")
	fmt.Scanln(&b)

	fmt.Print("enter the 3rd number: ")
	fmt.Scanln(&c)

	largest := a
	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}
	fmt.Println("largest number is:", largest)

}
