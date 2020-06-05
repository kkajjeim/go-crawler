package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"./src/watcha"
)

func main() {
	cpuNumber := 2
	runtime.GOMAXPROCS(cpuNumber)
	var wait sync.WaitGroup
	wait.Add(cpuNumber)

	startTime := time.Now()

	go func() {
		for i := 0; i < 15; i++ {
			watcha.Crawler()
		}

		defer wait.Done();wait.Done()
	}()

	go func() {
		for i := 0; i < 15; i++ {
			watcha.Crawler()
		}

		defer wait.Done()
	}()

	wait.Wait()
	elapsedTime := time.Since(startTime)

	fmt.Printf("duration: %s\n", elapsedTime)
}