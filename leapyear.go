package main

import "fmt"

func leapyear() {
	var year int
	fmt.Print("enter the year ")
	fmt.Scanln(&year)
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("It is  a leap year")
	} else {
		fmt.Println("It is not a leap year")
	}

}
