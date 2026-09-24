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
// Help received from: Oliwier Jakubiec, Mykhailo Balaker
// Help given to:
// Created on 24/09/2026
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
