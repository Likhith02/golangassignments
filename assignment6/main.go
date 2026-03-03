package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks float64
}

func problem1() {
	fmt.Println("--- 6.1 Define a Student Struct ---")
	var s Student

	fmt.Print("Enter name: ")
	fmt.Scan(&s.name)

	fmt.Print("Enter age: ")
	fmt.Scan(&s.age)

	fmt.Print("Enter marks: ")
	fmt.Scan(&s.marks)

	fmt.Println("Student Struct Created Successfully!")
	fmt.Printf("Name: %s, Age: %d, Marks: %.2f\n", s.name, s.age, s.marks)
}

func problem2() {
	fmt.Println("--- 6.2 Create and Print Student ---")
	var s Student

	fmt.Print("Enter name: ")
	fmt.Scan(&s.name)

	fmt.Print("Enter age: ")
	fmt.Scan(&s.age)

	fmt.Print("Enter marks: ")
	fmt.Scan(&s.marks)

	fmt.Println("\nStudent Details:")
	fmt.Printf("Name  : %s\n", s.name)
	fmt.Printf("Age   : %d\n", s.age)
	fmt.Printf("Marks : %.2f\n", s.marks)
}

func problem3() {
	fmt.Println("--- 6.3 Slice of Students ---")
	var n int
	fmt.Print("How many students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Student %d:\n", i+1)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].name)
		fmt.Print("Age: ")
		fmt.Scan(&students[i].age)
		fmt.Print("Marks: ")
		fmt.Scan(&students[i].marks)
	}

	fmt.Println("\nAll Students:")
	fmt.Println("-----------------------------")
	for i, s := range students {
		fmt.Printf("Student %d -> Name: %s, Age: %d, Marks: %.2f\n", i+1, s.name, s.age, s.marks)
	}
}

func problem4() {
	fmt.Println("--- 6.4 Student with Highest Marks ---")
	var n int
	fmt.Print("How many students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Student %d:\n", i+1)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].name)
		fmt.Print("Age: ")
		fmt.Scan(&students[i].age)
		fmt.Print("Marks: ")
		fmt.Scan(&students[i].marks)
	}

	top := students[0]
	for _, s := range students {
		if s.marks > top.marks {
			top = s
		}
	}

	fmt.Println("\nStudent with Highest Marks:")
	fmt.Println("-----------------------------")
	fmt.Printf("Name  : %s\n", top.name)
	fmt.Printf("Age   : %d\n", top.age)
	fmt.Printf("Marks : %.2f\n", top.marks)
}

func problem5() {
	fmt.Println("--- 6.5 Update Student Marks ---")
	var n int
	fmt.Print("How many students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Student %d:\n", i+1)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].name)
		fmt.Print("Age: ")
		fmt.Scan(&students[i].age)
		fmt.Print("Marks: ")
		fmt.Scan(&students[i].marks)
	}

	fmt.Println("\nAll Students:")
	for i, s := range students {
		fmt.Printf("%d. %s - %.2f marks\n", i+1, s.name, s.marks)
	}

	var index int
	fmt.Print("\nEnter student number to update marks (1 to ", n, "): ")
	fmt.Scan(&index)

	if index < 1 || index > n {
		fmt.Println("Invalid student number!")
		return
	}

	var newMarks float64
	fmt.Printf("Enter new marks for %s: ", students[index-1].name)
	fmt.Scan(&newMarks)

	fmt.Printf("\nBefore update: %s has %.2f marks\n", students[index-1].name, students[index-1].marks)
	students[index-1].marks = newMarks
	fmt.Printf("After update : %s has %.2f marks\n", students[index-1].name, students[index-1].marks)
}

func main() {
	var choice int
	fmt.Println("     ASSIGNMENT 6 - STRUCTS  ")
	fmt.Println("1. Define a Student Struct")
	fmt.Println("2. Create and Print a Student")
	fmt.Println("3. Slice of Students")
	fmt.Println("4. Student with Highest Marks")
	fmt.Println("5. Update Student Marks")
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
		fmt.Println("Invalid choice! Please enter a number between 1 and 5")
	}
}
