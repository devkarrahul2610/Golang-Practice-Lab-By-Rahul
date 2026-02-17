package main

import "fmt"

func AdvanceSlice() {
	fmt.Println("Welcome to slices Demo...")

	var cources = []string{"Java", "Go", "Docker", "Kubernaties"}
	fmt.Println("cources:", cources)
	fmt.Println("len:", len(cources), "cap:", cap(cources))
	cources = append(cources, "MySql")
	fmt.Println("cources:", cources)
	fmt.Println("len:", len(cources), "cap:", cap(cources))
}

func AdvanceSlice2() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}

func AdvanceSlice3() {
	base := make([]int, 0, 5)

	a := append(base, 1, 2)
	b := append(a, 3)
	c := append(a, 4)

	fmt.Println(base)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func update(s []int) {
	s = append(s, 100)
	s[0] = 999
}
