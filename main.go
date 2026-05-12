package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const (
	jobNumber    = 100
	workerNumber = 5
)

type Job struct {
	id       int
	duration int
}

func main() {
	jobs := make(chan Job, jobNumber)
	results := make(chan Job, jobNumber)
	count := 0

	for w := 0; w < workerNumber; w++ {
		go worker(w+1, jobs, results)
	}

	for j := 0; j < jobNumber; j++ {
		jobs <- generateJob(j + 1)
	}
	close(jobs)

	for a := 0; a < jobNumber; a++ {
		count++
		<-results
	}

	fmt.Printf("All jobs completed. Total jobs queued: %d. Total jobs done: %d\n", jobNumber, count)
}

func generateJob(id int) Job {
	return Job{
		id:       id,
		duration: rand.IntN(10) + 1,
	}
}

func worker(id int, jobs <-chan Job, results chan<- Job) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d with duration %d milliseconds\n", id, job.id, job.duration*10)
		time.Sleep(time.Duration(job.duration*10) * time.Millisecond)
		fmt.Printf("Worker %d finished job %d\n", id, job.id)
		results <- job
	}
}
