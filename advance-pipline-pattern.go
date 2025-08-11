package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//pattern 3: Pipline pattern. In go the pipline pattern is series of stages where each stage has one or more goroutines,
// reads from an input channel,performs some processing,send results to the output channel.
// stages are connected in a chain. The output of one is the input of the next.

func AdvancePiplinePattern() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//stage 1: Generator
	numCh := generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	//stage 2: fanout square stage with 3 workers.
	squaredCh := fanOut(ctx, numCh, 3)

	doubledCh := doubledStage(ctx, squaredCh)

	for result := range doubledCh {
		fmt.Println("Final Result :", result)
	}
}

// generator emits integers into channels until context is DONE.
func generator(ctx context.Context, val ...int64) <-chan int64 {

	outCh := make(chan int64)
	go func() {
		defer close(outCh)
		for _, v := range val {
			select {
			case <-ctx.Done():
				return
			case outCh <- v:
			}
		}

	}()
	return outCh
}

// creates a multiple workers for a stage.
func fanOut(ctx context.Context, inCh <-chan int64, workerCount int) chan int64 {
	outCh := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go squareWorker(ctx, i, inCh, outCh, &wg)
	}

	//close our channel when all workers are done.

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

func squareWorker(ctx context.Context, id int, inputCh <-chan int64, outputCh chan int64, wg *sync.WaitGroup) {

	defer wg.Done()
	for n := range inputCh {
		select {
		case <-ctx.Done():
			return
		case outputCh <- n * n:
			fmt.Printf("square worker %d processed %d\n", id, n)
			time.Sleep(time.Millisecond * time.Duration(rand.Int64N(300)))
		}

	}
}

func doubledStage(ctx context.Context, inCh chan int64) <-chan int64 {

	outCh := make(chan int64)

	go func() {
		defer close(outCh)
		for v := range inCh {
			select {
			case <-ctx.Done():
				return
			case outCh <- v * 2:
				fmt.Printf("Doubled %d---> %d\n", v, v*2)
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()
	return outCh

}
