package main

import (
	"fmt"
	"sync"
)

func advanceOfSlice1() {

	array := []int{10, 5, 12, 15, 20, 25, 18}

	s := array[1:5]

	s[1] = 23
	fmt.Println(s[1])
	fmt.Println(array)
}

/*  Write a Go program that implements a basic SlidingWindow processor using slices.
You must:
:one: Given a big slice of N integers (e.g., 1 to 100), create a sliding window of size k (e.g., 5).
 :two: Slide the window across the slice and, at each step, compute the average of the window.
 :three: Store all averages in a new slice.
 :four: While doing this, reuse slices efficiently — avoid unnecessary allocations.
 :five: If possible, use copy() smartly to show safe data handling when multiple goroutines are processing different windows (simulate with go func).
 :six: Finally, print:
the original slice
each window
each window’s average
the final list of averages */

func advanceOfSlice2() {

	N := 20
	K := 5

	// created a data slice
	data := make([]int, N)

	for i := 0; i < N; i++ {
		data[i] = i + 1
	}

	fmt.Println("Original Slice :", data)

	// how many windows there will be?
	numOfWindows := N - K + 1

	//Preallocate result slice of averages
	averages := make([]float64, numOfWindows)

	// use waitgroup to run safe gorutines for each window.
	var wg sync.WaitGroup
	wg.Add(numOfWindows)

	for i := 0; i < numOfWindows; i++ {

		// Get window view
		window := data[i : i+K]
		//fmt.Printf("Window %d: %v | add: %p\n", i, window, unsafe.Pointer(&window[0]))

		//If we pass window directly to goroutine, it may change next loop iteration.
		// so make a safe copy.

		safeWindow := make([]int, len(window))

		copy(safeWindow, window)

		go func(idx int, w []int) {
			defer wg.Done()
			avg := average(w)
			averages[idx] = avg

		}(i, safeWindow)

	}
	wg.Wait()
	fmt.Println("Final Averages :", averages)
}

func average(s []int) float64 {
	sum := 0
	for _, v := range s {
		sum = +v
	}

	return float64(sum) / float64(len(s))

}
