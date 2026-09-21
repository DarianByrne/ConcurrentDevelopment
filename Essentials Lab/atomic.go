// Author: Darian Byrne
// License: GPL3
// Help received from: Mykhailo Balaker

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup // IDE is NOT happy: 'wg' redeclared in this package

func addsAtomic(n int, total *atomic.Int64) bool {
	for i := 0; i < n; i++ {
		total.Add(1)
	}
	wg.Done()

	return true
}

func main() {
	var total atomic.Int64

	for i := range 10 {
		wg.Add(1)
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total)
	}
	wg.Wait()
	fmt.Println(total.Load()) // 10000
}
