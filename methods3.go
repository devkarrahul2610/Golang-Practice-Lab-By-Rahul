package main

import "fmt"

type Rectangle struct {
	length int64
	width  int64
}

func (r Rectangle) Area() int64 {
	return r.length * r.width
}

func Demo4() {
	var rectangle = Rectangle{length: 5, width: 6}

	area := rectangle.Area()
	fmt.Println("Area :", area)
}
