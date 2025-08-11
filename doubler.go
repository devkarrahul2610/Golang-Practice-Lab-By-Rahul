package main

import "sync"

func Generator(val ...int64) <-chan int64 {
	generatorCh := make(chan int64)

	go func() {
		defer close(generatorCh)
		for _, v := range val {
			generatorCh <- v
		}
	}()

	return generatorCh

}

func Doubler(ch <-chan int64) <-chan int64 {
	doublerCh := make(chan int64)

	go func() {
		defer close(doublerCh)
		for v := range ch {
			doublerCh <- v * 2
		}
	}()

	return doublerCh
}

func Merge(inch ...<-chan int64) <-chan int64 {
	outCh := make(chan int64)
	var wg sync.WaitGroup
	go func() {
		for _, ch := range inch {
			wg.Add(1)
			go func(c <-chan int64) {
				defer wg.Done()
				for v := range c {
					outCh <- v * 2
				}
			}(ch)
		}
		wg.Wait()
	}()
	return outCh
}
