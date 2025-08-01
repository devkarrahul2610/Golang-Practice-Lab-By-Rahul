package main

import "fmt"

// Example which proves maps are reference types in Go.
func advanceMap1() {

	Original := make(map[string]int)

	Original["one"] = 1

	fmt.Println("Values before modified :", Original)

	modify(Original)

	fmt.Println("Values after modified:", Original)

}

func modify(m map[string]int) {
	m["one"] = 10
}

// another example which proves maps are reference type in Go.

func advanceMap2() {
	Original := make(map[string]int)

	Original["one"] = 1

	fmt.Println("Values before modified :", Original)

	Copied := Original

	Copied["one"] = 10

	fmt.Println("Value after modified :", Original)
}

/* Assignment: Student Marks Manager
Scenario
You are building a system to:
Add students and their marks in multiple subjects.
Update marks for a subject.
Delete a student.
Calculate the average marks for each student.*/

func advanceMap3() {

	//create a map for student and their marks in multiple subjects

	students := make(map[string]map[string]int)

	// Add students and their marks in multiple subjects.
	addStudentMarks(students)
	fmt.Println("Students marks :", students)

	// update students marks
	updateStudentsMarks(students, "Rahul", "Science", 90)

	// delete student
	deleteStudent(students, "Rahul")
	fmt.Println("Students Record :", students)

	// calulate average marks for each students

	for name, _ := range students {
		calculateAverage(students, name)
	}

}
func calculateAverage(students map[string]map[string]int, name string) float64 {

	// check student is exist or not first.
	subjects, ok := students[name]
	if !ok {
		return 0.0 // student not found
	}
	sum := 0
	count := 0

	for _, marks := range subjects {
		sum += marks
		count++
	}

	if count == 0 {
		return 0.0
	}
	return float64(sum) / float64(count)
}

func deleteStudent(students map[string]map[string]int, name string) {

	if _, ok := students[name]; ok {
		delete(students, name)
	} else {
		fmt.Println("record not found")
	}
}

func updateStudentsMarks(students map[string]map[string]int, name, subject string, marks int) {

	if subjects, ok := students[name]; ok {
		subjects[subject] = marks
	} else {
		students[name] = map[string]int{subject: marks}
	}
}

func addStudentMarks(s map[string]map[string]int) {

	s["Rahul"] = map[string]int{
		"Maths":   90,
		"Science": 85,
	}

	s["Pratiksha"] = map[string]int{
		"Maths":   98,
		"History": 94,
		"Science": 88,
	}

	s["Aishu"] = map[string]int{
		"Maths":   96,
		"History": 93,
	}
}
