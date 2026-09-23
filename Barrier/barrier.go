package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func doStuff(goNum int, wg *sync.WaitGroup) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	// barrier here
	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	sem.Acquire(ctx, 1)
	for i := range totalRoutines {
		go doStuff(i, &wg)
	}
	sem.Release(1)
	theLock.Unlock()

	wg.Wait()
}
