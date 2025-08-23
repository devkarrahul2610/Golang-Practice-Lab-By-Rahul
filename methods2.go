package main

import (
	"fmt"
	"strings"
)

// FullName is a custom function type.
// It represents any function that takes two strings and returns a string.
type FullName func(string, string) string

// UpperCaseForFuncType is a method on the FullName type.
// ✅ Yes, in Go you can attach methods to any named type, including function types.
//
// How it works:
// - `f FullName` is the receiver, so this method can be called on variables of type FullName.
// - Inside the method, we invoke the underlying function `f(firstName, lastName)`.
// - Then we convert the result to uppercase.
func (f FullName) UpperCaseForFuncType(firstName string, lastName string) string {
	return strings.ToUpper(string(f(firstName, lastName)))
}

func Demo3() {
	// Define a variable of type FullName.
	// Assign an anonymous function (lambda) that matches the signature: func(string, string) string.
	var fullname FullName = func(s1, s2 string) string { // Anonymous function assigned
		return s1 + "****" + s2
	}

	// Directly call the function variable
	fmt.Println(fullname("rahul", "devkar"))
	// Output: rahul****devkar

	// Call the method UpperCaseForFuncType on the function variable
	// Internally it calls fullname("rahul", "devkar"), then converts it to uppercase.
	fmt.Println(fullname.UpperCaseForFuncType("rahul", "devkar"))
	// Output: RAHUL****DEVKAR
}
