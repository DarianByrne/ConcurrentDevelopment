package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep((X * time.Second))
	fmt.Println("Part A", Num)
	// rendezvous here
	fmt.Println("Part B", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N)
	}
	wg.Wait()
}
