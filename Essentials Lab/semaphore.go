// Author: Darian Byrne
// License: GPL3
// Help received from: Seamus Kennedy

package main

import (
	"fmt"
	"sync"
	"time"
)

// this type is unused? so i commented it out to make the IDE happy
//type semaphore struct {
//	theCounter chan struct{}
//}

// funcAcquite(sem *Semaphore) // what is this?
func main() {
	maxGoRoutines := 5
	semaphore := make(chan struct{}, maxGoRoutines)

	var wg sync.WaitGroup
	for i := range 20 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
