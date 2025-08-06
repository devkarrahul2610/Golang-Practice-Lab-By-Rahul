package main

import "fmt"

func arrayBasic() {

	var arr [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println(arr)

	a := [5]int{5, 10, 12, 4, 18}

	var p *[5]int = &a

	for i := 0; i < len(a); i++ {
		fmt.Println((*p)[i])

	}

	fmt.Println(a)

	name := [2]string{"Rahul", "Devkar"}
	fmt.Println(name)

	var names = [2]string{"Rahul", "Devkar"}
	fmt.Println(names)
}

func arrayBasicOne() {
	var arr = [5]int{5, 10, 15, 20, 25}

	fmt.Println(arr)

	fmt.Println(&arr[0])
	fmt.Println(&arr[1])

	var p *[5]int = &arr

	fmt.Println((p)[0])
}

func arrayBasictwo() {
	var arr = [5]int{5, 10, 15, 20, 25}
	var p *[5]int = &arr

	for i := 0; i < len(arr); i++ {

		fmt.Println(&(p)[i])
	}
}

func arrayBasicThree() {
	// take the input from the user

	var arr [5]int

	fmt.Println("Take the input from the user")

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element %d :", i+1)
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func arrayBasicFour() {
	// prove how array is copy by value

	var arr = [5]int{5, 10, 15, 20, 25}

	fmt.Println("Before Modifying :", arr)

	modify(arr)

	fmt.Println("After Modifying :", arr)

}

func modify(a [5]int) {
	// here the array values are copied,Its not the same array which we declared on arrayBasicFour()
	a[1] = 100
	fmt.Println(a)
}
