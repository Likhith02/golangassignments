package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scanln(&n)
	a := 0
	b := 1
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
}
