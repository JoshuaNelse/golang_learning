package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		myLoop("test1")
		waitGroup.Done()
	}()

	myLoop("test2")
	waitGroup.Wait()
}

func myLoop(variable string) {
	for i := 0; i < 5; i++ {
		fmt.Println(variable)
	}
}
