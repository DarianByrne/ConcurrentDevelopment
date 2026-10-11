package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func reader(wg *sync.WaitGroup, rMutex *sync.Mutex, wMutex *sync.Mutex, count *int, num int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	rMutex.Lock()
	if *count == 0 {
		wMutex.Lock() // don't allow writers to write while reader reads, unless there's a writer already writing
	}
	*count++
	rMutex.Unlock()

	// critical section
	fmt.Println(num, "reading")

	rMutex.Lock()
	*count--
	if *count == 0 {
		wMutex.Unlock()
	}
	rMutex.Unlock()
	wg.Done()
}

func writer(wg *sync.WaitGroup, theLock *sync.Mutex, num int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	theLock.Lock()

	// critical section
	fmt.Println(num, "writing...")
	time.Sleep(3 * time.Second)
	fmt.Println(num, "finished writing")

	theLock.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var rMutex sync.Mutex
	var wMutex sync.Mutex
	numReaders := 10
	numWriters := 10
	wg.Add(numReaders + numWriters)
	readerCount := 0

	for i := range numReaders {
		go reader(&wg, &rMutex, &wMutex, &readerCount, i)
	}
	for i := range numWriters {
		go writer(&wg, &wMutex, i)
	}

	wg.Wait() //wait here until everyone is done
}
