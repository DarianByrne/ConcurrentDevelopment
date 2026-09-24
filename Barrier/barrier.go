//rendezvous.go
//Copyright (C) 2026 Darian Byrne

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Darian Byrne
// Help received from: Joseph Kehoe, Mykhailo Balaker, Oliwier Jakubiec
// Help given to:
// Created on 23/09/2026
// Modified by:
// Issues:
//--------------------------------------------

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
