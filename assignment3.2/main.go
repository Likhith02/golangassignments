package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	if isPrime(num) {
		fmt.Println("It is a Prime number")
	} else {
		fmt.Println("It is not a Prime Number")
	}
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
