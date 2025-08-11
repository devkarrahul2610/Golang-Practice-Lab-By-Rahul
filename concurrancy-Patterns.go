package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// There are several well established pattern using channels in go. These are used in real life example to build scalable
// concurrent piplines.

// 1. FanOut/FanIn patterns:
// UseCase: Distribute works to multiple go-routines(FanOut) and Collects results into single channel(FanIn).

func Worker(id int, jobs chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d Processing Job %d\n", id, job)
		results <- job * 2
	}

}

// Worker Pool Pattern:
//UseCase: Control Concurrancy Level. Avoid overwhelming resources.
// worker pool is a concurrancy pattern where you have a fixed number of worker goroutines. These worked reads task from a common channel.
// This allows bounded concurrancy: avoiding spanning up too many gorotines and exhausting system resources.

// processes job from channel
func Worker2(id int, jobs chan Job, result chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		//simulate work
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		output := job.RandomNo * 2

		result <- Result{JobID: job.ID, Input: job.RandomNo, Output: output, WorkerID: id}

	}
}
