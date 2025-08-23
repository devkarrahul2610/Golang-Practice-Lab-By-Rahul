package main

import (
	"errors"
	"fmt"
)

// Define a new named type 'marks'.
// Underlying type is: map[string]int
// - string → student name
// - int    → marks scored
type marks map[string]int

// add is a method on the 'marks' type.
// It adds a new entry (student → marks) into the map.
//
// Receiver: (m marks)
//   - Since 'marks' is a map type (reference type), any modifications inside
//     the method affect the original map (no need for a pointer receiver).
//
// Logic:
//  1. Check if the key already exists.
//  2. If yes → return an error saying "key is already present".
//  3. If not → insert the new key-value pair into the map.
func (m marks) add(k string, val int) error {
	if _, ok := m[k]; ok {
		return errors.New("key is already present: " + k)
	}
	m[k] = val
	return nil
}

func Demo2() {
	// Create a variable of type 'marks' using make.
	// 'make' initializes the map so it's ready for use.
	var marksData = make(marks)

	// Try adding a new student
	if err := marksData.add("Rahul", 90); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("****** marks added successfully ******")
	}

	// Try adding the same student again (will trigger error path)
	if err := marksData.add("Rahul", 85); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("****** marks added successfully ******")
	}

	// Print the final map
	fmt.Println("Final marks data:", marksData)
}
