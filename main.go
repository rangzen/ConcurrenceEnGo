package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	func() {
		compte("mouton")
		wg.Done()
	}()
	wg.Wait()
}

func compte(chose string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, chose)
		time.Sleep(500 * time.Millisecond)
	}
}
