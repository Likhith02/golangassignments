package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks float64
}

func double(n *int) {
	*n = *n * 2
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func updateStudent(s *Student, newName string, newAge int, newMarks float64) {
	s.name = newName
	s.age = newAge
	s.marks = newMarks
}

func problem1() {
	fmt.Println("--- 8.1 Pointer Demonstration ---")

	var x int
	fmt.Print("Enter a number: ")
	fmt.Scan(&x)

	p := &x

	fmt.Println("\n--- Pointer Details ---")
	fmt.Printf("Value of x        : %d\n", x)
	fmt.Printf("Address of x (&x) : %v\n", &x)
	fmt.Printf("Pointer p holds   : %v\n", p)
	fmt.Printf("Value via *p      : %d\n", *p)

	var newVal int
	fmt.Print("\nEnter new value to set via pointer: ")
	fmt.Scan(&newVal)

	*p = newVal

	fmt.Println("\n--- After Modifying via Pointer ---")
	fmt.Printf("Value of x now    : %d\n", x)
	fmt.Printf("Value via *p now  : %d\n", *p)
}

func problem2() {
	fmt.Println("--- 8.2 Modify Variable Using Pointer ---")

	var num int
	fmt.Print("Enter a number to double: ")
	fmt.Scan(&num)

	fmt.Printf("\nBefore double() : %d\n", num)
	double(&num)
	fmt.Printf("After double()  : %d\n", num)
}

func problem3() {
	fmt.Println("--- 8.3 Swap Two Numbers ---")

	var a, b int
	fmt.Print("Enter first number : ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Printf("\nBefore swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap : a = %d, b = %d\n", a, b)
}

func problem4() {
	fmt.Println("--- 8.4 Modify Struct Using Pointer ---")

	var s Student

	fmt.Println("Enter original student details:")
	fmt.Print("Name  : ")
	fmt.Scan(&s.name)
	fmt.Print("Age   : ")
	fmt.Scan(&s.age)
	fmt.Print("Marks : ")
	fmt.Scan(&s.marks)

	fmt.Println("\n--- Before Update ---")
	fmt.Printf("Name  : %s\n", s.name)
	fmt.Printf("Age   : %d\n", s.age)
	fmt.Printf("Marks : %.2f\n", s.marks)

	var newName string
	var newAge int
	var newMarks float64

	fmt.Println("\nEnter new student details:")
	fmt.Print("New Name  : ")
	fmt.Scan(&newName)
	fmt.Print("New Age   : ")
	fmt.Scan(&newAge)
	fmt.Print("New Marks : ")
	fmt.Scan(&newMarks)

	updateStudent(&s, newName, newAge, newMarks)

	fmt.Println("\n--- After Update ---")
	fmt.Printf("Name  : %s\n", s.name)
	fmt.Printf("Age   : %d\n", s.age)
	fmt.Printf("Marks : %.2f\n", s.marks)
}

func main() {
	var choice int

	fmt.Println("    ASSIGNMENT 8 - POINTERS  ")
	fmt.Println("1. Pointer Demonstration")
	fmt.Println("2. Modify Variable Using Pointer")
	fmt.Println("3. Swap Two Numbers")
	fmt.Println("4. Modify Struct Using Pointer")
	fmt.Print("Enter problem number (1-4): ")
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
	default:
		fmt.Println("Invalid choice! Please enter a number between 1 and 4")
	}
}
