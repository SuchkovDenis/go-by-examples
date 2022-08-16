package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println("worker", id, "started job", i)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", i)
		results <- i * 2
	}
}

func main() {
	const numJobs = 15
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
