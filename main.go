package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func log() {
	fmt.Println("log")
}

func main() {
	wg.Add(1)

	go func() {
		defer wg.Done()
		log()
	}()

	wg.Wait()

	fmt.Println("done")
}
