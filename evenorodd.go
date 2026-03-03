package main

import "fmt"

func evenorodd() {

	var num int
	fmt.Print("Enter a number ")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

}
