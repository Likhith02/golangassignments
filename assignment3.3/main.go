package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	reversed := reverseNumber(num)
	fmt.Println("Reversed number is:", reversed)

}
func reverseNumber(n int) int {
	reverse := 0
	for n != 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n = n / 10

	}
	return reverse

}
