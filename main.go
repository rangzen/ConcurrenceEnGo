package main

import (
	"fmt"
	"time"
)

func main() {
	go compte("mouton")
	go compte("poisson")

	fmt.Scanln()
}

func compte(chose string) {
	for i := 1; true; i++ {
		fmt.Println(i, chose)
		time.Sleep(500 * time.Millisecond)
	}
}
