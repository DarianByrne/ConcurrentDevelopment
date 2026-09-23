// Author: Darian Byrne
// License: GPL3
// Help received from: Oliwier Jakubiec, Mykhailo Balaker

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, c1 *chan struct{}, c2 *chan struct{}) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep((X * time.Second))
	fmt.Println("Part A", Num)

	// rendezvous here
	if Num == 0 {
		*c2 <- struct{}{}
		<-*c1
	} else {
		<-*c2
		*c1 <- struct{}{}
	}

	fmt.Println("Part B", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 2

	c1 := make(chan struct{}, threadCount)
	c2 := make(chan struct{}, threadCount)

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &c1, &c2)
	}
	wg.Wait()
}
