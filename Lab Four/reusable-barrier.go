//reusable-barrier.go
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
// Help received from: Mykhailo Balaker, Oliwier Jakubiec
// Help given to:
// Created on 28/09/2026
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

// Create a barrier data type
type barrier struct {
	theChan chan bool
	theLock sync.Mutex
	total   int
	count   int
}

// creates a properly initialised barrier
// N== number of threads (go Routines)
func createBarrier(N int) barrier {
	theBarrier := barrier{
		theChan: make(chan bool),
		total:   N,
		count:   0,
	}
	return theBarrier
}

// Method belonging to barrier data type
// Blocks until everyone reaches this point then lets everyone continue
func (b *barrier) wait() {
	b.theLock.Lock()
	b.count++
	if b.count == b.total {
		b.count = 0
		b.theLock.Unlock()
		fmt.Println("here")
		for _ = range b.total - 1 {
			<-b.theChan
		}
	} else {
		fmt.Println(b.count)
		b.theLock.Unlock()
		b.theChan <- true
	}
} //wait

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, barrierOne *barrier, barrierTwo *barrier) bool {
	for N := range 3 {
		var X time.Duration
		X = time.Duration(rand.IntN(5))
		time.Sleep(X * time.Second) //wait random time amount
		fmt.Println("Part A", Num, ", Loop: ", N)
		//Rendezvous here
		barrierOne.wait()
		fmt.Println("Part B", Num, ", Loop: ", N)
		barrierTwo.wait()
	}
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrierA := createBarrier(5)
	barrierB := createBarrier(5)
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &barrierA, &barrierB)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done

}
