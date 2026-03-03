package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks float64
}

func (s Student) display() {
	fmt.Println("\nStudent Details:")
	fmt.Printf("Name  : %s\n", s.name)
	fmt.Printf("Age   : %d\n", s.age)
	fmt.Printf("Marks : %.2f\n", s.marks)
}

func (s *Student) updateMarks(newMarks float64) {
	s.marks = newMarks
}

func (s Student) hasPassed() bool {
	return s.marks >= 40
}

func (s Student) grade() string {
	switch {
	case s.marks >= 90:
		return "A+"
	case s.marks >= 80:
		return "A"
	case s.marks >= 70:
		return "B"
	case s.marks >= 60:
		return "C"
	case s.marks >= 40:
		return "D"
	default:
		return "F"
	}
}

func problem1() {
	fmt.Println("--- 7.1 Display Student Details ---")
	var s Student

	fmt.Print("Enter name: ")
	fmt.Scan(&s.name)
	fmt.Print("Enter age: ")
	fmt.Scan(&s.age)
	fmt.Print("Enter marks: ")
	fmt.Scan(&s.marks)

	s.display()
}

func problem2() {
	fmt.Println("--- 7.2 Update Student Marks ---")
	var s Student

	fmt.Print("Enter name: ")
	fmt.Scan(&s.name)
	fmt.Print("Enter age: ")
	fmt.Scan(&s.age)
	fmt.Print("Enter marks: ")
	fmt.Scan(&s.marks)

	fmt.Printf("\nCurrent marks of %s: %.2f\n", s.name, s.marks)

	var newMarks float64
	fmt.Print("Enter new marks: ")
	fmt.Scan(&newMarks)

	s.updateMarks(newMarks)

	fmt.Printf("Updated marks of %s: %.2f\n", s.name, s.marks)
}

func problem3() {
	fmt.Println("--- 7.3 Check Pass or Fail ---")
	var s Student

	fmt.Print("Enter name: ")
	fmt.Scan(&s.name)
	fmt.Print("Enter age: ")
	fmt.Scan(&s.age)
	fmt.Print("Enter marks: ")
	fmt.Scan(&s.marks)

	s.display()

	if s.hasPassed() {
		fmt.Printf("\n%s has PASSED! ✓\n", s.name)
	} else {
		fmt.Printf("\n%s has FAILED! ✗\n", s.name)
	}
}

func problem4() {
	fmt.Println("--- 7.4 Get Student Grade ---")
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

	fmt.Println("\nStudent Results:")
	fmt.Println("-----------------------------")
	for _, s := range students {
		fmt.Printf("Name: %-10s | Marks: %.2f | Grade: %s | Status: ", s.name, s.marks, s.grade())
		if s.hasPassed() {
			fmt.Println("PASSED ✓")
		} else {
			fmt.Println("FAILED ✗")
		}
	}
}

func main() {
	var choice int

	fmt.Println("    ASSIGNMENT 7 - METHODS   ")
	fmt.Println("1. Display Student Details")
	fmt.Println("2. Update Student Marks")
	fmt.Println("3. Check Pass or Fail")
	fmt.Println("4. Get Student Grade")
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
