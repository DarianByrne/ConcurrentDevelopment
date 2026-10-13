package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func producer(theChan *chan int) {
	for {
		var X time.Duration
		X = time.Duration(rand.IntN(5))
		time.Sleep(X * time.Second) //wait random time amount

		val := rand.IntN(50)
		fmt.Println("Produced", val)
		*theChan <- val
	}
}

func consumer(theChan *chan int) {
	for {
		var value int
		value = <-*theChan
		fmt.Println("Consumed", value)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	theChan := make(chan int)
	go producer(&theChan)
	go consumer(&theChan)
	wg.Wait()
}
