package main

import (
	"fmt"
	"strings"
)

// why would we use type?
// if var already works whats the point of creating new type?
// Ans: 1. Type safety: we can prevent accidental mixing of values that are semantically different, even if they share same underlying
// type. eg.

func TypeDemo1() {

	type Kilometer float64
	type miles float64

	// var d1 Kilometer = 5
	// var d2 miles = 5

	// d1 = d2 // will get compile error (different types)
	// This ensures we don't accidently assigne miles to Kilometers.
}

// 2. we can't attach methods to build in types(eg. string, int). But if we define new type then we can.
type Name string //define a new type from string(existing type)

// NewUpperCase is a method not just a regular function. Why? because it has a receiver(n Name).
// In go function of a receiver is called a method of that type.
func (n Name) NewUpperCase() Name {
	return Name(strings.ToUpper(string(n)))
}

func MethodDemo() {

	var MyName Name

	MyName = "RahulDevkar"

	upperLetters := MyName.NewUpperCase()
	fmt.Println(upperLetters)

}

//3. only types can implement interface.
// By defining your own type, you can make it satisfy an interface.
// fmt.Stringer interface.

func (n Name) String() string {
	return "My name is :" + string(n)
}
func InterfaceDemo() {

	var myName Name

	myName = "RahulSanjayDevkar"

	fmt.Println(myName)
}
