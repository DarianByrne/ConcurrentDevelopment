// Author: Darian Byrne
// License: GPL3
// Help received from: Joseph Kehoe, Mykhailo Balaker, Oliwier Jakubiec

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func doStuff(goNum int, wg *sync.WaitGroup, theLock *sync.Mutex, theChan *chan struct{}, total int, count *int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	fmt.Println("Part A", goNum)

	// barrier here
	theLock.Lock() // must lock before updating/checking count, otherwise another thread could update it
	*count++
	if *count == total { // if the last thread
		theLock.Unlock()
		*theChan <- struct{}{} // start releasing everyone
	} else {
		theLock.Unlock()
		<-*theChan             // not the last thread so must wait
		*theChan <- struct{}{} // i'm released, now i must release someone else
	}
	theLock.Lock() // must lock before updating/checking count
	*count--
	if *count == 0 {
		<-*theChan // last thread closes
	}
	theLock.Unlock()

	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	totalRoutines := 10

	count := 0
	var barrierLock sync.Mutex
	barrierChannel := make(chan struct{}, totalRoutines)

	wg.Add(totalRoutines)
	for i := range totalRoutines {
		go doStuff(i, &wg, &barrierLock, &barrierChannel, totalRoutines, &count)
	}
	wg.Wait()
}
