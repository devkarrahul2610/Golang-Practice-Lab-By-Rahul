package main

import (
	"fmt"
	"strings"
)

// this is a normal function
func UpperCase(n string) {
	fmt.Println(strings.ToUpper(n))
}

func Demo1() {

	var name string

	name = "RahulDevkar"

	// name:="RahulDevkar"  short declaration.

	UpperCase(name)
}
