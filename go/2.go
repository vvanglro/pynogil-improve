package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func cpuBoundTask(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i * i
	}
	return sum
}

func main() {
	threads := flag.Int("threads", 4, "Number of threads")
	tasks := flag.Int("tasks", 10, "Number of tasks")
	size := flag.Int("size", 5000000, "Task size (n for sum of squares)")
	flag.Parse()

	fmt.Printf("Running %d tasks of size %d with %d threads\n", *tasks, *size, *threads)
	runtime.GOMAXPROCS(*threads)
	startTime := time.Now()

	done := make(chan bool)

	for _ = range *tasks {
		go func() {
			cpuBoundTask(*size)
			done <- true
		}()
	}

	for i := 0; i < *tasks; i++ {
		<-done
	}
	endTime := time.Now()
	fmt.Println(endTime)
	duration := endTime.Sub(startTime)

	fmt.Printf("Time with goroutines: %v seconds\n", duration.Seconds())
}
