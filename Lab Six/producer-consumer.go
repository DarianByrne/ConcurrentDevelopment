package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func producer(theChan *chan int, num int) {
	for {
		var X time.Duration
		X = time.Duration(rand.IntN(5))
		time.Sleep(X * time.Second) //wait random time amount

		val := rand.IntN(100)
		fmt.Println(num, "Produced", val, "after waiting", int(X), "second(s)")
		*theChan <- val
	}
}

func consumer(theChan *chan int, num int) {
	for {
		var value int
		value = <-*theChan
		fmt.Println(num, "Consumed", value)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	theChan := make(chan int)
	numProducers := 10
	numConsumers := 2
	for i := range numProducers {
		go producer(&theChan, i)
	}
	for i := range numConsumers {
		go consumer(&theChan, i)
	}
	wg.Wait()
}
