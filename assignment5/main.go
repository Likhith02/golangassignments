package main

import (
	"fmt"
	"sort"
)

func problem1() {
	fmt.Println("1. Create a slice of integers and print all elements")
	var n int
	fmt.Print("How many numbers you want to enter:")
	fmt.Scan(&n)
	nums := make([]int, n)
	fmt.Println("Enter", n, "integers:")
	for i := 0; i < n; i++ {
		fmt.Printf("number %d", i+1)
		fmt.Scan(&nums[i])
	}
	fmt.Println("slice numbers:")
	for _, v := range nums {
		fmt.Println(v)
	}
}
func problem2() {
	fmt.Println("2.Add elements to a slice using append")
	var n int
	fmt.Print("How many initial elements: ")
	fmt.Scan(&n)
	nums := make([]int, n)
	fmt.Println("enter", n, "integers:")
	for i := 0; i < n; i++ {
		fmt.Printf("element %d:", i+1)
		fmt.Scan(&nums[i])
	}
	fmt.Println("before append:", nums)
	var m int
	fmt.Print("How many elements to append:")
	fmt.Scan(&m)
	fmt.Println("Enter", m, "integers to append:")
	for i := 0; i < m; i++ {
		var val int
		fmt.Printf("element %d: ", i+1)
		fmt.Scan(&val)
		nums = append(nums, val)
	}
	fmt.Println("After append:", nums)
}
func problem3() {
	fmt.Println("3.Remove an element from a slice at a given index")
	var n int
	fmt.Print("How many elements:")
	fmt.Scan(&n)
	nums := make([]int, n)
	fmt.Println("Enter", n, "integers: ")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d:", i+1)
		fmt.Scan(&nums[i])

	}
	fmt.Println("Slice:", nums)
	var index int
	fmt.Printf("Enter index to remove (0 to %d):", n-1)
	fmt.Scan(&index)
	if index < 0 || index >= n {
		fmt.Println("Invalid index")
		return
	}
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println("After removing index", index, ":", nums)

}

func problem4() {
	fmt.Println("4.Find the sum and average of elements in a slice.")
	var n int
	fmt.Print("How many elemtns:")
	fmt.Scan(&n)
	nums := make([]int, n)
	fmt.Println("Enter", n, "integers:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&nums[i])
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	avg := float64(sum) / float64(len(nums))
	fmt.Println("Slice:", nums)
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", avg)
}
func problem5() {
	fmt.Println("5.Sort a slice in ascending order.")
	var n int
	fmt.Print("How many elements:")
	fmt.Scan(&n)
	nums := make([]int, n)
	fmt.Println("enter", n, "integers:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&nums[i])

	}
	fmt.Println("before sort:", nums)
	sort.Ints(nums)
	fmt.Println("After sort:", nums)
}

func main() {
	var choice int
	fmt.Println("1. Create a slice of integers and print all elements.\n2. Add elements to a slice using append.\n3. Remove an element from a slice at a given index.\n4. Find the sum and average of elements in a slice.\n5. Sort a slice in ascending order.")
	fmt.Print("Enter problem number (1-5): ")
	fmt.Scan(&choice)
	fmt.Println()
	switch choice {
	case 1:
		problem1()
	case 2:
		problem2()
	case 3:
		problem3()
	case 4:
		problem4()
	case 5:
		problem5()
	default:
		fmt.Println("wrong choice please enter a number from 1 to 5")
	}

}
