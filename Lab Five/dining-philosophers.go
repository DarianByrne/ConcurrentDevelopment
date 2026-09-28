//dining-philosophers.go
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

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

func getForks(index int, forks map[int]chan bool) {
	forks[index] <- true
	forks[(index+1)%5] <- true
}

func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //main
