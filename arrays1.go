package main

import (
	"fmt"
)

func advanceArrayOne() {
	// two dimentional array

	arr := [2][2]int{{5, 10}, {15, 20}}
	fmt.Println("len :", len(arr))

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			//fmt.Println(arr[i][j])
			fmt.Printf("arr[%d][%d] : %d\n", i, j, arr[i][j])
		}
	}

}

func sortArraysAssending() {

	arr := [5]int{55, 10, 17, 12, 9}

	var temp int

	fmt.Println("Before Sorting :", arr)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	fmt.Println("After Sorting :", arr)
}

func rotateArray() {

	//not a recommended approach
	arr := [5]int{10, 4, 8, 12, 93}

	mArr := [5]int{}

	index := 2
	j := 0
	k := 0
	for i := index + 1; i < len(arr); i++ {
		mArr[j] = arr[i]
		j++
	}

	for i := index; i < len(arr); i++ {
		mArr[i] = arr[k]
		k++

	}
	fmt.Println(mArr)

}

func rotateArrayNew() {

	arr := [5]int{10, 4, 8, 12, 93}
	var rotatedArray [5]int
	// var n int
	// fmt.Println("Took the index from user")
	// fmt.Scanf("index : %d", &n)
	n := 2
	for i := 0; i < len(arr); i++ {
		rotatedArray[i] = arr[(i+n)%(len(arr))]
	}

	fmt.Println(rotatedArray)
}
