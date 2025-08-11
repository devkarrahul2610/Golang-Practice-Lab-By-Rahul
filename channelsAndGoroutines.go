package main

import (
	"fmt"
	"time"
)

func cook(dish string) {

	for i := 1; i <= 3; i++ {
		fmt.Println("Cooking", dish, "step", i)
		time.Sleep(time.Second)
	}

}

func unbufferedChannel() {
	ch := make(chan int64)

	ch <- 10 // unbuffered stuck, no space on channel to put this value.

	val := <-ch
	fmt.Println("ch :", val)
}

func unbufferedChannel2() {
	ch := make(chan int64)

	go func() {
		defer close(ch)
		ch <- 10 // understand like this one goroutine is created for writing the values on channel
	}()

	val := <-ch // here we are reading this value from the channel. But understand here main function is also the goroutines and unbufferedChannel2()
	// is the function called from main. SO its like we are running another goroutines for reading the value.
	fmt.Println("val :", val)

	//val = <-ch
	//fmt.Println("val :", val) // there is no value present inside the channel and if we are trying to read this it will be on deadlock.

}

func bufferedChannel() {
	ch := make(chan int64, 2) // Here we have created a space for 2 values inside the channels. So we can put directly two values
	// inside the channel
	defer close(ch)
	ch <- 10
	ch <- 20

	val := <-ch
	fmt.Println("val :", val)

	val = <-ch
	fmt.Println("val :", val)
}
