// Author: Darian Byrne
// License: GPL3
// Help received from: Milosz Cwynar

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")

		barrier <- true

		fmt.Println("StuffOne - Part B")
		wg.Done()
		return true
	}

	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")

		<-barrier

		fmt.Println("StuffTwo - Part B")
		wg.Done()
		return true
	}

	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()
	wg.Wait()
}
