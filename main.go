package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// In Golang, go routine is light-weight,user-space thread managed by go runtime. It is extremly cheap to create starting with just
// 2KB of stack trace, and scales dynamically. You start it using the go keyword.
// ex. go doTask()
// behind the scenes what happed here? Ans--> Go doesn't create a new OS thread for each goroutine. Instead it uses its own GMP schedular
// to multiplex thousands of goroutines over a small number of OS threads.
// In Go, concurrancy means managing multiple task at once (Interleaved Execution) while parrelelisum means execute multiple task
// simulteniously on multiple CPU cores.
// By default Go is concurrent,But with multiple cores (GOMAXPROS), it can be parellel too.
// runtime.GOMAXPROS(runtime.NumCPU) enables true parellelisum

// Channels are typed conduits through which goroutines communicate. They enforce synchronization and make data sharing safe without
// explicit locks.
func main1() {
	fmt.Println("Welcome to the Go Advance Tutorial")
	//unbufferedChannel()
	unbufferedChannel2()
	fmt.Println("End of the execution")
	time.Sleep(time.Second)
}

func main2() {
	fmt.Println("Started Execution Here")
	go cook("PIZZA") // runs concurrently
	go cook("PASTA")
	fmt.Println("Execution ending HERE")
	time.Sleep(time.Second * 4) // waiting for finish go-routines
}

func main3() {
	fmt.Println("Welcome to the Go Advance Tutorial")
	bufferedChannel()
	fmt.Println("End of the execution")
}

func main4() {
	fmt.Println("Welcome to the Double Demo")

	ch := Generator(1, 3, 4, 7, 9, 5, 11)
	outch := Doubler(ch)

	for v := range outch {
		fmt.Println("Double values are :", v)
	}
}

// Doubler-Merge Assignment. Pipline Channel Pattern.
func main5() {

	generatorCh := Generator(1, 3, 4, 7, 9, 5, 11)

	dCh1 := Doubler(generatorCh)
	dCh2 := Doubler(generatorCh)
	dCh3 := Doubler(generatorCh)

	outCh := Merge(dCh1, dCh2, dCh3)
	fmt.Println("Double Values :", outCh)
}

// 1. FanOut/FanIn
func main6() {
	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go Worker(w, jobCh, resultCh, &wg)
	}

	// send jobs

	for i := 1; i <= 5; i++ {
		jobCh <- i
	}
	close(jobCh)
	wg.Wait()
	close(resultCh)

	//FanIn: collect the result

	for result := range resultCh {
		fmt.Println("Result :", result)
	}

}

//2. Worker Pool pattern.

// Job represents a number of work
type Job struct {
	ID       int
	RandomNo int
}

// Result holds the result of processing job.
type Result struct {
	JobID    int
	Input    int
	Output   int
	WorkerID int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numbJobs := 10
	numWorkers := 3

	jobs := make(chan Job, numbJobs)
	results := make(chan Result, numbJobs)

	var wg sync.WaitGroup

	// start worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Worker2(w, jobs, results, &wg)
	}

	//send jobs

	for j := 1; j <= numbJobs; j++ {
		jobs <- Job{ID: j, RandomNo: rand.Intn(100)}
	}
	close(jobs)

	// close result channels after all workers finish

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Worker %d processed Job %d: input=%d output=%d\n",
			result.WorkerID, result.JobID, result.Input, result.Output)
	}
}
