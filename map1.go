package main

import (
	"fmt"
	"strings"
)

func basicMap1() {
	// Declare the map
	var myMap map[string]int

	// Initialize the map
	myMap = make(map[string]int)

	//Direct Initialization of map
	//myNewMap := make(map[string]int)

	myMap["rahul"] = 27
	myMap["pratiksha"] = 22
	fmt.Println("map :", myMap)
}

// counting words
func basicMap2() {

	text := "apple banana apple orange banana apple"

	words := make(map[string]int)

	for _, word := range strings.Split(text, " ") {
		words[word]++
	}

	fmt.Println(words)
}

// If you want to read something from map

func basicMap3() {
	infoMap := make(map[string]int)

	infoMap["rahul"] = 27
	infoMap["pratiksha"] = 22
	infoMap["aishu"] = 25

	// how to read something from map

	age := infoMap["aishu"]

	fmt.Println("aishu age :", age)
}

// check if data is exist.
func basicMap4() {
	infoMap := make(map[string]int)

	infoMap["rahul"] = 27
	infoMap["pratiksha"] = 22
	infoMap["aishu"] = 25
	val, ok := infoMap["pratiksha"]
	if ok {
		fmt.Println("val :", val)
	} else {
		fmt.Println("Data Not found")
	}
}

// how to delete any data from map
func basicMap5() {
	infoMap := make(map[string]int)

	infoMap["rahul"] = 27
	infoMap["pratiksha"] = 22
	infoMap["aishu"] = 25
	_, ok := infoMap["rahul"]
	if ok {
		delete(infoMap, "rahul")
	} else {
		fmt.Println("data not found.")
	}
	fmt.Println(infoMap)
}

// Loop through maps
func basicMaps6() {

	infoMap := make(map[string]int)

	infoMap["rahul"] = 27
	infoMap["pratiksha"] = 22
	infoMap["aishu"] = 25

	for key, value := range infoMap {
		fmt.Println("key :", key+" "+"val:", value)
	}
}
