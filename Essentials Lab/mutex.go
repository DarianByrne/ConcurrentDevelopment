// Author: Darian Byrne
// License: GPL3
// Help given to: Thomas Radulescu, Filip Raguz

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var total int64

func adds(n int, theLock *sync.Mutex) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		total++
		theLock.Unlock()
	}
	wg.Done()
	return true
}

func main() {
	var theLock sync.Mutex

	total = 0
	wg.Add(10)

	for i := range 10 {
		fmt.Println(i)
		go adds(1000, &theLock)
	}
	wg.Wait()
	fmt.Println(total) // 10000
}
