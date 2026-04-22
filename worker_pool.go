package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id int
}

type Result struct {
	task   Task
	status string
}

type Stats struct {
	totalProcessed int
	mutex          sync.Mutex
}

func (s *Stats) incrementTotal() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.totalProcessed++
}

func startWorkerPool() {
	fmt.Println("WORKER POOL STARTING...")
	tasksChannel := make(chan Task, 20)
	resultsChannel := make(chan Result, 20)
	stats := Stats{}

	for x := 0; x < 4; x++ {
		go processTask(&stats, tasksChannel, resultsChannel)
	}

	for x := 0; x < 20; x++ {
		tasksChannel <- Task{id: x}
		fmt.Printf("send task to channel %d\n", x)
	}

	// sender should close channel as soon as sending is done
	close(tasksChannel)

	for x := 0; x < 20; x++ {
		result := <-resultsChannel
		fmt.Printf("recieve result from process %d\n", result.task.id)
	}

	fmt.Printf("Total processed: %d\n", stats.totalProcessed)

	fmt.Printf("WORKER POOL FINISHED.\n\n")
}

func processTask(stats *Stats, tasksChannel <-chan Task, resultsChannel chan Result) {
	for task := range tasksChannel {
		// do task
		time.Sleep(time.Millisecond * 150)

		stats.incrementTotal()
		resultsChannel <- Result{task: task}
	}
}
