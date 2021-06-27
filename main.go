package main

import (
	"fmt"
	"time"
)

func main() {
	compte("mouton")
	compte("poisson")
}

func compte(chose string) {
	for i := 1; true; i++ {
		fmt.Println(i, chose)
		time.Sleep(500 * time.Millisecond)
	}
}
